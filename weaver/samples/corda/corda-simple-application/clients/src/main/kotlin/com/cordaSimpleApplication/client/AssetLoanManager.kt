/*
 * Copyright IBM Corp. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package com.cordaSimpleApplication.client

import arrow.core.Either
import arrow.core.Left
import arrow.core.Right
import com.github.ajalt.clikt.core.CliktCommand
import com.github.ajalt.clikt.core.requireObject
import com.github.ajalt.clikt.parameters.arguments.argument
import com.github.ajalt.clikt.parameters.options.option
import com.github.ajalt.clikt.parameters.options.flag
import com.github.ajalt.clikt.parameters.options.default
import java.io.File
import java.lang.Exception
import kotlinx.coroutines.runBlocking
import com.google.protobuf.util.JsonFormat
import java.util.Base64
import java.time.Instant
import kotlin.system.exitProcess
import com.google.gson.Gson
import com.google.protobuf.ByteString
import java.util.Calendar
// Corda
import net.corda.core.utilities.OpaqueBytes
import net.corda.core.crypto.sha256
import net.corda.core.identity.CordaX500Name
import net.corda.core.identity.Party
import net.corda.core.messaging.startFlow
import net.corda.core.messaging.CordaRPCOps
// cacti weaver sdk
import org.hyperledger.cacti.weaver.sdk.corda.HashFunctions
import org.hyperledger.cacti.weaver.sdk.corda.AssetTransferSDK
// cacti weaver imodule
import org.hyperledger.cacti.weaver.imodule.corda.states.AssetClaimHTLCData
import org.hyperledger.cacti.weaver.imodule.corda.states.AssetPledgeState
import org.hyperledger.cacti.weaver.imodule.corda.flows.getHashMechanism
import org.hyperledger.cacti.weaver.imodule.corda.flows.RetrieveNetworkId
// sample cordapp
import com.cordaSimpleApplication.state.AssetState
import com.cordaSimpleApplication.state.LoanRepaymentCondition
import com.cordaSimpleApplication.contract.AssetContract
import com.cordaSimpleApplication.flow.ClaimAndPledgeAssetStateInitiator
import com.cordaSimpleApplication.flow.ClaimLoanedAssetInitiator
import com.cordaSimpleApplication.flow.ClaimLoanRepaymentInitiator

object AssetLoanManager {
    class LoanCommand : CliktCommand(name = "loan", help ="Manages simple asset loan") {
        override fun run() {
        }
    }

    /**
     * Claim and Pledge the asset
     */
    class ClaimAndPledgeCommand : CliktCommand(
            help = "Claim And Pledge an asset. claim-and-pledge --secret=secret --loan_period=10 --recipient='PartyA'") {
        val config by requireObject<Map<String, String>>()
        val contractId: String? by option("-cid", "--contract-id", help="Contract/Linear Id for HTLC State")
        val hash_fn: String? by option("-hfn", "--hash-fn", help="Hash Function to be used. Default: SHA256")
        val secret: String? by option("-s", "--secret", help="Hash Pre-Image for the HTLC Claim")
        val token_ledger: String? by option("-tlid", "--token-ledger", help="Token Ledger Id")
        val token_type: String? by option("-tt", "--token-type", help="Token Type")
        val repayment_amount: String? by option("-ra", "--repayment-amount", help="Amount of tokens for repayment")
        val token_lender: String? by option("-tl", "--token-lender", help="Borrower name in Token ledger as per remoteNetworkConfig.json")
        val token_borrower: String? by option("-tb", "--token-borrower", help="Borrower name in Token ledger as per remoteNetworkConfig.json")
        val loan_period: String? by option("-l", "--loan-period", help="Asset Loan period in days")
        val observer: String? by option("-o", "--observer", help="Party Name for Observer")
        override fun run() = runBlocking {
            var hash: HashFunctions.Hash = HashFunctions.SHA256()
            if(hash_fn == "SHA256") {
                hash = HashFunctions.SHA256()
            } else if ( hash_fn == "SHA512") {
                hash = HashFunctions.SHA512()
            }
            if (contractId == null || secret == null) {
                println("Arguments required: --contract-id and --secret.")
            } else {
                val rpc = NodeRPCConnection(
                        host = config["CORDA_HOST"]!!,
                        username = "clientUser1",
                        password = "test",
                        rpcPort = config["CORDA_PORT"]!!.toInt())
                try {
                    val issuer = rpc.proxy.wellKnownPartyFromX500Name(CordaX500Name.parse("O=PartyA,L=London,C=GB"))!!
                    hash.setPreimage(secret!!)
                    val claimInfoData = AssetClaimHTLCData(
                        hashMechanism = getHashMechanism(hash.HASH_MECHANISM),
                        hashPreimage = OpaqueBytes(hash.getPreimage()!!.toByteArray())
                    )
                    val tokenLedgerBorrowerCert: String = getUserCertFromFile(token_borrower!!, token_ledger!!)
                    val tokenLedgerLenderCert: String = getUserCertFromFile(token_lender!!, token_ledger!!)
                    val pledgeCondition = LoanRepaymentCondition(
                        tokenType = token_type!!,
                        tokenQuantity = repayment_amount!!.toLong(),
                        tokenLedgerId = token_ledger!!,
                        tokenLedgerLenderCert = tokenLedgerLenderCert,
                        tokenLedgerBorrowerCert = tokenLedgerBorrowerCert,
                        assetType = "",
                        assetId = "",
                        assetLedgerId = "",
                        assetLedgerLenderCert = "",
                        assetLedgerBorrowerCert = ""
                    )
                    var obs = listOf<Party>()
                    if (observer != null)   {
                        obs += rpc.proxy.wellKnownPartyFromX500Name(CordaX500Name.parse(observer!!))!!
                    }
                    val result = runCatching {
                        rpc.proxy.startFlow(::ClaimAndPledgeAssetStateInitiator, contractId!!, claimInfoData, pledgeCondition, loan_period!!.toLong() * 24L * 3600L, issuer, obs)
                            .returnValue.get()
                    }.fold({ it ->
                        it.map { pledgeId ->
                            println("Claim and Pledging asset was successful and the pledge state was stored with pledgeId $pledgeId.\n")
                            pledgeId.toString()
                        }
                    }, { it -> 
                        println("Corda Network Error: Error running ClaimAndPledgeAssetState flow: ${it.message}\n")
                    })
                    println("Asset ClaimAndPledgeAssetState Response: ${result}")
                } catch (e: Exception) {
                  println("Error: ${e.toString()}")
                } finally {
                    rpc.close()
                }
            }
        }
    }
    
    /**
     * Command to pledge an asset.
     * transfer pledge-asset --timeout=120 -rnid 'Corda_Network2' -nrid 'Corda_Network2' --recipient='<name of the recipient>' --param=type:id ----> non-fungible
     * transfer pledge-asset --fungible --timeout=120 -rnid Corda_Network2 --recipient='<name of the recipient>' --param=type:amount ----> fungible
     */
    class PledgeTokensCommand : CliktCommand(name="pledge-asset",
            help = "Locks an asset. $ ./clients transfer pledge-asset --fungible --timeout=10 -rnid 'Corda_Network2' --recipient='<name of recipient>' --param=type:amount") {
        val config by requireObject<Map<String, String>>()
        val timeout: String? by option("-t", "--timeout", help="Pledge validity time duration in seconds.")
        val importNetworkId: String? by option("-inid", "--import-network-id", help="Importing network for asset transfer")
        val recipient: String? by option("-r", "--recipient", help="Name of the recipient in the importing network")
        val param: String? by option("-p", "--param", help="Parameter AssetType:AssetId for non-fungible, AssetType:Quantity for fungible.")
        val observer: String? by option("-o", "--observer", help="Party Name for Observer")
        override fun run() = runBlocking {
            if (recipient == null) {
                println("Arguement -r (name of the recipient in importing n/w) is required")
            } else if (param == null) {
                println("Arguement -p (asset details to be pledged) is required")
            } else if (importNetworkId == null) {
                println("Arguement -inid (importing/remote network id) is required")
            } else {
                var nTimeout: Long
                if (timeout == null) {
                    nTimeout = 300L
                } else {
                    nTimeout = timeout!!.toLong()
                }
                val calendar = Calendar.getInstance()
                nTimeout += calendar.timeInMillis / 1000
                println("nTimeout: $nTimeout")

                val rpc = NodeRPCConnection(
                        host = config["CORDA_HOST"]!!,
                        username = "clientUser1",
                        password = "test",
                        rpcPort = config["CORDA_PORT"]!!.toInt())
                try {
                    val params = param!!.split(":").toTypedArray()
                    
                    
                    val localNetworkId = rpc.proxy.startFlow(::RetrieveNetworkId).returnValue.get()
                    println("localNetworkId: ${localNetworkId}")

                     // "thisParty" is set to the token "issuer" in case fungible house token; since we are using the same
                     // SDK function claimPledgeFungibleAsset and Interop application for both the "Simple Asset" and
                     // the "Fungible house token" corDapps, we pass the Identity of the party submitting the claim here.
                    val issuer: Party = rpc.proxy.wellKnownPartyFromX500Name(CordaX500Name.parse(ISSUER_DN))!!

                    var obs = listOf<Party>()
                    if (observer != null)   {
                       obs += rpc.proxy.wellKnownPartyFromX500Name(CordaX500Name.parse(observer!!))!!
                    }

                    // Obtain the recipient certificate from the name of the recipient
                    val recipientCert: String = getUserCertFromFile(recipient!!, importNetworkId!!)


                    val result = AssetTransferSDK.createFungibleAssetPledge(
                        rpc.proxy,
                        localNetworkId!!,
                        importNetworkId!!,
                        params[0],          // Type
                        params[1].toLong(), // Quantity
                        recipientCert,
                        nTimeout,
                        "com.cordaSimpleApplication.flow.RetrieveStateAndRef",
                        AssetContract.Commands.LoanPledge(),
                        issuer,
                        obs
                    )
                    
                    
                    when (result) {
                        is Either.Left -> {
                            println("Corda Network Error: Error running PledgeAsset flow: ${result.a.message}\n")
                            throw IllegalStateException("Corda Network Error: Error running PledgeAsset flow: ${result.a.message}\n")
                        }
                        is Either.Right -> {
                            println("AssetPledgeState created with pledge-id '${result.b}'")
                        }
                    }
                } catch (e: Exception) {
                  println("Error: ${e.toString()}")
                } finally {
                    rpc.close()
                }
            }
        }
    }
    /**
     * Claim the loaned asset
     */
    class LoanClaimAssetCommand : CliktCommand(
            help = "Claim a loaned asset. claim-asset --pledge-id=abc --recipient='PartyA'") {
        val config by requireObject<Map<String, String>>()
        val pledgeId: String? by option("-pid", "--pledge-id", help="Pledge id for asset loan pledge state")
        val remotePledgeId: String? by option("-rpid", "--remote-pledge-id", help="Pledge id for token pledged state")
        val lender: String? by option("-l", "--lender", help="X500 name for lender Party in this asset ledger")
        val assetRelayAddress: String? by option ("-ar", "--asset-relay", help="This (asset) ledger relay address")
        val tokenLedgerType: String? by option("-tlt", "--token-ledger-type", help="DLT type of remote network: fabric|corda|besu")
        val tokenLedgerId: String? by option("-tlid", "--token-ledger-id", help="Ledger id for token network")
        val observer: String? by option("-o", "--observer", help="Party Name for Observer")
        override fun run() = runBlocking {
            if (pledgeId == null ||
                remotePledgeId == null ||
                lender == null ||
                assetRelayAddress == null ||
                tokenLedgerType == null ||
                tokenLedgerId == null) {
                println("Arguments required: --pledge-id, --remote-pledge-id, --lender, --asset-ledger-id, --token-ledger-id, --token-ledger-type, and --asset-relay.")
            } else {
                val rpc = NodeRPCConnection(
                        host = config["CORDA_HOST"]!!,
                        username = "clientUser1",
                        password = "test",
                        rpcPort = config["CORDA_PORT"]!!.toInt())
                try {
                    var assetPledgeState: AssetPledgeState
                    when (val result: Either<Error, AssetPledgeState> = AssetTransferSDK.getAssetPledgeStatus(rpc.proxy, pledgeId!!, tokenLedgerId!!)) {
                        is Either.Left -> {
                            println("Corda Network Error: Error running GetAssetPledgeStatus flow: ${result.a.message}\n")
                            throw IllegalStateException("Corda Network Error: Error running GetAssetPledgeStatus flow: ${result.a.message}\n")
                        }
                        is Either.Right -> {
                            assetPledgeState = result.b
                        }
                    }
                    // "thisParty" is set to the token "issuer" in case fungible house token; since we are using the same
                    // SDK function claimPledgeFungibleAsset and Interop application for both the "Simple Asset" and
                    // the "Fungible house token" corDapps, we pass the Identity of the party submitting the claim here.
                    val issuer = rpc.proxy.wellKnownPartyFromX500Name(CordaX500Name.parse("O=PartyA,L=London,C=GB"))!!
                    var obs = listOf<Party>()
                    if (observer != null)   {
                        obs += rpc.proxy.wellKnownPartyFromX500Name(CordaX500Name.parse(observer!!))!!
                    }

                    // Obtain the locker certificate from the name of the locker
                    val assetLedgerLenderCert = assetPledgeState.lockerCert
                    val tokenLedgerBorrowerCert = assetPledgeState.recipientCert

                    var externalStateAddress: String = getClaimLoanedAssetViewAddress(
                        tokenLedgerType!!, 
                        remotePledgeId!!, 
                        assetPledgeState.localNetworkId, 
                        assetPledgeState.remoteNetworkId, 
                        assetLedgerLenderCert, 
                        tokenLedgerBorrowerCert
                    )

                    // 1. While exercising 'data transfer' initiated by a Corda network, the localRelayAddress is obtained directly from user.
                    // 2. While exercising 'asset transfer' initiated by a Fabric network, the localRelayAddress is obtained from config.json file
                    // 3. While exercising 'asset transfer' initiated by a Corda network (this case), the localRelayAddress is obtained
                    //    below from the remote-network-config.json file
                    //val networkConfig: JSONObject = getRemoteNetworkConfig(importNetworkId)
                    //val importRelayAddress: String = networkConfig.getString("relayEndpoint")
                    val pledgeStatusLinearId: String = requestStateFromRemoteNetwork(assetRelayAddress!!, externalStateAddress, rpc.proxy, config, listOf(issuer))

                    val result = runCatching {
                        rpc.proxy.startFlow(::ClaimLoanedAssetInitiator, pledgeId!!, pledgeStatusLinearId, issuer, obs)
                            .returnValue.get()
                    }.fold({ it ->
                        it.map { pledgeId ->
                            println("Claim loaned asset was successful and claim status was stored with id $pledgeId.\n")
                            pledgeId.toString()
                        }
                    }, { it -> 
                        println("Corda Network Error: Error running ClaimLoanedAsset flow: ${it.message}\n")
                    })
                    println("Loaned asset claim by borrower response: ${result}")
                } catch (e: Exception) {
                    println("Error: ${e.toString()}")
                    // exit the process throwing error code
                    exitProcess(1)
                } finally {
                    rpc.close()
                }
            }
        }
    }
    /**
     * Claim repayment
     */
    class LoanClaimRepaymentCommand : CliktCommand(
            help = "Claim a loaned payment. claim-repayment --pledge-id=abc --recipient='PartyA'") {
        val config by requireObject<Map<String, String>>()
        val pledgeId: String? by option("-pid", "--pledge-id", help="Pledge id for asset loan pledge state")
        val remotePledgeId: String? by option("-rpid", "--remote-pledge-id", help="Pledge id for token pledged state")
        val borrower: String? by option("-b", "--borrower", help="X500 name for borrower Party in this asset ledger")
        val tokenRelayAddress: String? by option ("-tr", "--token-relay", help="This (token) ledger relay address")
        val assetLedgerType: String? by option("-tlt", "--asset-ledger-type", help="DLT type of remote network: fabric|corda|besu")
        val assetLedgerId: String? by option("-tlid", "--asset-ledger-id", help="Ledger ID for asset network")
        val observer: String? by option("-o", "--observer", help="Party Name for Observer")
        override fun run() = runBlocking {
            if (pledgeId == null ||
                remotePledgeId == null ||
                borrower == null ||
                tokenRelayAddress == null ||
                assetLedgerType == null ||
                assetLedgerId == null) {
                println("Arguments required: --pledge-id, --remote-pledge-id, --borrower, --token-ledger-id, --asset-ledger-id, --asset-ledger-type, and --asset-relay.")
            } else {
                val rpc = NodeRPCConnection(
                        host = config["CORDA_HOST"]!!,
                        username = "clientUser1",
                        password = "test",
                        rpcPort = config["CORDA_PORT"]!!.toInt())
                try {
                    var assetPledgeState: AssetPledgeState
                    when (val result: Either<Error, AssetPledgeState> = AssetTransferSDK.getAssetPledgeStatus(rpc.proxy, pledgeId!!, assetLedgerId!!)) {
                        is Either.Left -> {
                            println("Corda Network Error: Error running GetAssetPledgeStatus flow: ${result.a.message}\n")
                            throw IllegalStateException("Corda Network Error: Error running GetAssetPledgeStatus flow: ${result.a.message}\n")
                        }
                        is Either.Right -> {
                            assetPledgeState = result.b
                        }
                    }
                    // "thisParty" is set to the token "issuer" in case fungible house token; since we are using the same
                    // SDK function claimPledgeFungibleAsset and Interop application for both the "Simple Asset" and
                    // the "Fungible house token" corDapps, we pass the Identity of the party submitting the claim here.
                    val issuer = rpc.proxy.wellKnownPartyFromX500Name(CordaX500Name.parse("O=PartyA,L=London,C=GB"))!!
                    var obs = listOf<Party>()
                    if (observer != null)   {
                        obs += rpc.proxy.wellKnownPartyFromX500Name(CordaX500Name.parse(observer!!))!!
                    }

                    // Obtain the locker certificate from the name of the locker
                    val tokenLedgerBorrowerCert = assetPledgeState.lockerCert
                    val assetLedgerLenderCert = assetPledgeState.recipientCert
                    //val lenderCert: String = getUserCertFromFile(lender!!, tokenLedgerId!!)
                    val borrower = rpc.proxy.wellKnownPartyFromX500Name(CordaX500Name.parse(borrower!!))!!

                    val pledgeCondition = Gson().fromJson(ByteString.copyFrom(assetPledgeState.pledgeCondition).toStringUtf8(), LoanRepaymentCondition::class.java)

                    var externalStateAddress: String = getClaimRepaymentViewAddress(
                        assetLedgerType!!, 
                        remotePledgeId!!, 
                        assetPledgeState.expiryTimeSecs.toString(), 
                        assetPledgeState.localNetworkId, 
                        assetPledgeState.remoteNetworkId,
                        pledgeCondition.assetType,
                        pledgeCondition.assetId,
                        tokenLedgerBorrowerCert, 
                        assetLedgerLenderCert
                    )

                    // 1. While exercising 'data transfer' initiated by a Corda network, the localRelayAddress is obtained directly from user.
                    // 2. While exercising 'asset transfer' initiated by a Fabric network, the localRelayAddress is obtained from config.json file
                    // 3. While exercising 'asset transfer' initiated by a Corda network (this case), the localRelayAddress is obtained
                    //    below from the remote-network-config.json file
                    //val networkConfig: JSONObject = getRemoteNetworkConfig(importNetworkId)
                    //val importRelayAddress: String = networkConfig.getString("relayEndpoint")
                    val claimStatusLinearId: String = requestStateFromRemoteNetwork(tokenRelayAddress!!, externalStateAddress, rpc.proxy, config, listOf(issuer))

                    val result = runCatching {
                        rpc.proxy.startFlow(::ClaimLoanRepaymentInitiator, pledgeId!!, claimStatusLinearId, issuer, obs)
                            .returnValue.get()
                    }.fold({ it ->
                        it.map { pledgeId ->
                            println("Claim loan repayment was successful and claim status was stored with id $pledgeId.\n")
                            pledgeId.toString()
                        }
                    }, { it -> 
                        println("Corda Network Error: Error running ClaimLoanRepaymentInitiator flow: ${it.message}\n")
                    })
                    println("Loan repayment claim by lender response: ${result}")
                } catch (e: Exception) {
                    println("Error: ${e.toString()}")
                    // exit the process throwing error code
                    exitProcess(1)
                } finally {
                    rpc.close()
                }
            }
        }
    }
}

fun getClaimLoanedAssetViewAddress(tokenLedgerType: String, remotePledgeId: String, localNetworkId: String, remoteNetworkId: String, assetLenderCert: String?, tokenBorrowerCert: String?): String {
    // assetLenderCert -> locker
    // tokenBorrowerCert -> recipient
    if (tokenLedgerType.equals("corda")) {
        return generateViewAddressFromRemoteConfig(remoteNetworkId, "GetAssetPledgeStatusByPledgeId", listOf(remotePledgeId, localNetworkId))
    } else if (tokenLedgerType.equals("fabric")) {
        throw Error("Unsupported ledger type: ${tokenLedgerType}")
        //return generateViewAddressFromRemoteConfig(remoteNetworkId, "GetAssetPledgeStatus", listOf(remotePledgeId, assetLenderCert!!, localNetworkId, tokenBorrowerCert!!))
    } else {
        throw Error("Unsupported ledger type: ${tokenLedgerType}")
    }
}
fun getClaimRepaymentViewAddress(assetLedgerType: String, remotePledgeId: String, expiryTimeSecs: String, localNetworkId: String, remoteNetworkId: String, assetType: String?, assetId: String?, tokenBorrowerCert: String?, assetLenderCert: String?): String {
    // tokenBorrowerCert -> locker
    // assetLenderCert -> recipient
    if (assetLedgerType.equals("corda")) {
        return generateViewAddressFromRemoteConfig(remoteNetworkId, "GetBondAssetClaimStatusByPledgeId", listOf(remotePledgeId, expiryTimeSecs))
    } else if (assetLedgerType.equals("fabric")) {
        throw Error("Unsupported ledger type: ${assetLedgerType}")
        //return generateViewAddressFromRemoteConfig(remoteNetworkId, "GetAssetClaimStatus", listOf(remotePledgeId, assetType!!, assetId!!, assetLenderCert!!, tokenBorrowerCert!!, remoteNetworkId, expiryTimeSecs))
    } else {
        throw Error("Unsupported ledger type: ${assetLedgerType}")
    }
}