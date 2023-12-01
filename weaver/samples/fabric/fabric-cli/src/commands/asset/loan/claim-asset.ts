/*
 * Copyright IBM Corp. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

import { GluegunCommand } from 'gluegun'
import * as path from 'path'
import {
    commandHelp,
    pledgeAsset,
    getNetworkConfig,
    getLocalAssetPledgeDetails,
    getUserCertFromFile,
    getChaincodeConfig,
    handlePromise,
    generateViewAddressFromRemoteConfig,
    interopHelper
} from '../../../helpers/helpers'
import {
    fabricHelper,
    getUserCertBase64
} from '../../../helpers/fabric-functions'
import { getLoanRepaymentCondition } from '../../../helpers/loan'

import logger from '../../../helpers/logger'
import * as dotenv from 'dotenv'
dotenv.config({ path: path.resolve(__dirname, '../../../.env') })

const delay = ms => new Promise(res => setTimeout(res, ms))

const command: GluegunCommand = {
  name: 'claim-asset',
  description:
    'Claims pledged asset in destination network',
  run: async toolbox => {
    const {
      print,
      parameters: { options, array }
    } = toolbox
    if (options.help || options.h) {
      commandHelp(
        print,
        toolbox,
        'fabric-cli asset loan claim-asset --token-network=network2 --asset-network=network1 --borrower=alice --lender=bob --type=bond --pledge-id="<pledgeid>" --token-pledge-id="<token-pledge-id>"',
        'fabric-cli asset loan claim-asset --token-network=<token-network-name> --asset-network=<asset-network-name> --borrower=<borrower-id> --type=<bond|token> --pledge-id=<pledge-id> --param=<asset-type>:<asset-id|num-units>',
        [
          {
            name: '--debug',
            description:
              'Shows debug logs when running. Disabled by default. To enable --debug=true'
          },
          {
            name: '--token-network',
            description:
              'Network where the asset is currently present. <network1|network2>'
          },
          {
            name: '--asset-network',
            description:
              'Network where the asset is to be transferred. <network1|network2>'
          },
          {
            name: '--lender',
            description:
              'Lender name'
          },
          {
            name: '--borrower',
            description:
              'Borrower name'
          },
          {
            name: '--pledge-id',
            description:
              'Pledge Id associated with asset loan.'
          },
          {
            name: '--token-pledge-id',
            description:
              'Pledge Id associated with asset loan.'
          },
          {
            name: '--relay-tls',
            description: 'Flag indicating whether or not the relay is TLS-enabled.'
          },
          {
            name: '--relay-tls-ca-files',
            description: 'Colon-separated list of root CA certificate paths used to connect to the relay over TLS.'
          },
          {
            name: '--e2e-confidentiality',
            description: 'Flag indicating whether or not the view contents are confidential end-to-end across networks (client-to-interop-module).'
          }
        ],
        command,
        ['asset', 'loan', 'claim-asset']
      )
      return
    }

    if (options.debug === 'true') {
      logger.level = 'debug'
      logger.debug('Debugging is enabled')
    }
    if (!options['asset-network'])
    {
      print.error('--asset-network needs to be specified')
      return
    }
    if (!options['token-network'])
    {
      print.error('--token-network needs to be specified')
      return
    }
    if (!options['lender'])
    {
      print.error('--lender needs to be specified')
      return
    }
    if (!options['borrower'])
    {
      print.error('--borrower needs to be specified')
      return
    }
    if (!options['type'])
    {
      print.error('--type of asset loan needs to be specified in the format: \'asset_type.remote_network_type\'.' +
            ' \'asset_type\' can be either \'bond\', \'token\' or \'house-token\'.' +
            ' \'remote_network_type\' can be either \'fabric\', \'corda\' or \'besu\'.')
      return
    }
    if (!options['pledge-id'])
    {
      print.error('--pledge-id needs to be specified')
      return
    }
    if (!options['token-pledge-id'])
    {
      print.error('--token-pledge-id needs to be specified')
      return
    }
    
    
    const netConfig = getNetworkConfig(options['asset-network'])
    if (!netConfig.connProfilePath || !netConfig.channelName || !netConfig.chaincode) {
      print.error(
        `Please use a valid --asset-network. No valid environment found for ${options['asset-network']} `
      )
      return
    }
    const transferCategory = 'token.fabric'
    
    try {
      const borrowerCert = await getUserCertBase64(options['asset-network'], options['borrower'])
      const { viewAddress, ownerCert } = await getClaimViewAddress(transferCategory, options['token-pledge-id'],
        options['borrower'], options['token-network'], borrowerCert, options['asset-network']
      )
      
      const applicationFunction = 'ClaimLoanedAsset'
      var { args, replaceIndices } = getChaincodeConfig(netConfig.chaincode, applicationFunction)
      args[args.indexOf('<pledge-id>')] = options['pledge-id']
      args[args.indexOf('token-network')] = options['token-network']
      options['user'] = options['borrower'] 
      await interopHelper(
        options['asset-network'],
        viewAddress,
        netConfig.chaincode,
        applicationFunction,
        args,
        replaceIndices,
        options,
        print        
      )
      process.exit()
    } catch (error) {
      print.error(`Error Asset Loan Claim: ${error}`)
    }
  }
}


async function getClaimViewAddress(transferCategory, pledgeId, owner, sourceNetwork,
    recipientCert, destNetwork
) {
    let funcName = "", funcArgs = []
    let ownerCert = await getUserCertFromFile(owner, sourceNetwork)

    if (transferCategory == "token.corda") {
        funcName = "GetAssetPledgeStatusByPledgeId"
        funcArgs = [pledgeId, destNetwork]
    } else if (transferCategory === "bond.fabric") {
        funcName = "GetAssetPledgeStatus"
        funcArgs = [pledgeId, ownerCert, destNetwork, recipientCert]
    } else if (transferCategory === "token.fabric") {
        funcName = "GetTokenAssetPledgeStatus"
        funcArgs = [pledgeId, ownerCert, destNetwork, recipientCert]
    } else if (transferCategory.includes("house-token.corda")) {
        funcName = "GetAssetPledgeStatusByPledgeId"
        funcArgs = [pledgeId, destNetwork]
    } else {
        throw new Error(`Unecognized loan category: ${transferCategory}`)
    }

    const viewAddress = generateViewAddressFromRemoteConfig(sourceNetwork, funcName, funcArgs)

    return { viewAddress: viewAddress, ownerCert: ownerCert }
}

module.exports = command

