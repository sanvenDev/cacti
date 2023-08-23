/**
 *
 * Please note:
 * This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * Do not edit this file manually.
 *
 */

@file:Suppress(
    "ArrayInDataClass",
    "EnumEntryName",
    "RemoveRedundantQualifierName",
    "UnusedImport"
)

package org.openapitools.client.models


import com.squareup.moshi.Json

/**
 * 
 *
 * @param key The key that was used to check the presence of the value in the entry store.
 * @param checkedAt Date and time encoded as JSON when the presence check was performed by the plugin backend.
 * @param isPresent The boolean true or false indicating the presence or absence of an entry under 'key'.
 */


data class HasKeychainEntryV1200Response (

    /* The key that was used to check the presence of the value in the entry store. */
    @Json(name = "key")
    val key: kotlin.Any?,

    /* Date and time encoded as JSON when the presence check was performed by the plugin backend. */
    @Json(name = "checkedAt")
    val checkedAt: kotlin.Any?,

    /* The boolean true or false indicating the presence or absence of an entry under 'key'. */
    @Json(name = "isPresent")
    val isPresent: kotlin.Any?

)

