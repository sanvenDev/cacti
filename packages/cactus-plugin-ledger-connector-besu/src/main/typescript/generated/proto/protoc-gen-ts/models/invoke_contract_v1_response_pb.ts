/**
 * Generated by the protoc-gen-ts.  DO NOT EDIT!
 * compiler version: 3.19.1
 * source: models/invoke_contract_v1_response_pb.proto
 * git: https://github.com/thesayyn/protoc-gen-ts */
import * as dependency_1 from "./../google/protobuf/any";
import * as dependency_2 from "./web3_transaction_receipt_pb";
import * as pb_1 from "google-protobuf";
export namespace org.hyperledger.cacti.plugin.ledger.connector.besu {
    export class InvokeContractV1ResponsePB extends pb_1.Message {
        #one_of_decls: number[][] = [[401260801]];
        constructor(data?: any[] | ({
            transactionReceipt?: dependency_2.org.hyperledger.cacti.plugin.ledger.connector.besu.Web3TransactionReceiptPB;
            success?: boolean;
        } & (({
            callOutput?: dependency_1.google.protobuf.Any;
        })))) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], this.#one_of_decls);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("transactionReceipt" in data && data.transactionReceipt != undefined) {
                    this.transactionReceipt = data.transactionReceipt;
                }
                if ("callOutput" in data && data.callOutput != undefined) {
                    this.callOutput = data.callOutput;
                }
                if ("success" in data && data.success != undefined) {
                    this.success = data.success;
                }
            }
        }
        get transactionReceipt() {
            return pb_1.Message.getWrapperField(this, dependency_2.org.hyperledger.cacti.plugin.ledger.connector.besu.Web3TransactionReceiptPB, 472834426) as dependency_2.org.hyperledger.cacti.plugin.ledger.connector.besu.Web3TransactionReceiptPB;
        }
        set transactionReceipt(value: dependency_2.org.hyperledger.cacti.plugin.ledger.connector.besu.Web3TransactionReceiptPB) {
            pb_1.Message.setWrapperField(this, 472834426, value);
        }
        get has_transactionReceipt() {
            return pb_1.Message.getField(this, 472834426) != null;
        }
        get callOutput() {
            return pb_1.Message.getWrapperField(this, dependency_1.google.protobuf.Any, 401260801) as dependency_1.google.protobuf.Any;
        }
        set callOutput(value: dependency_1.google.protobuf.Any) {
            pb_1.Message.setOneofWrapperField(this, 401260801, this.#one_of_decls[0], value);
        }
        get has_callOutput() {
            return pb_1.Message.getField(this, 401260801) != null;
        }
        get success() {
            return pb_1.Message.getFieldWithDefault(this, 256557056, false) as boolean;
        }
        set success(value: boolean) {
            pb_1.Message.setField(this, 256557056, value);
        }
        get _callOutput() {
            const cases: {
                [index: number]: "none" | "callOutput";
            } = {
                0: "none",
                401260801: "callOutput"
            };
            return cases[pb_1.Message.computeOneofCase(this, [401260801])];
        }
        static fromObject(data: {
            transactionReceipt?: ReturnType<typeof dependency_2.org.hyperledger.cacti.plugin.ledger.connector.besu.Web3TransactionReceiptPB.prototype.toObject>;
            callOutput?: ReturnType<typeof dependency_1.google.protobuf.Any.prototype.toObject>;
            success?: boolean;
        }): InvokeContractV1ResponsePB {
            const message = new InvokeContractV1ResponsePB({});
            if (data.transactionReceipt != null) {
                message.transactionReceipt = dependency_2.org.hyperledger.cacti.plugin.ledger.connector.besu.Web3TransactionReceiptPB.fromObject(data.transactionReceipt);
            }
            if (data.callOutput != null) {
                message.callOutput = dependency_1.google.protobuf.Any.fromObject(data.callOutput);
            }
            if (data.success != null) {
                message.success = data.success;
            }
            return message;
        }
        toObject() {
            const data: {
                transactionReceipt?: ReturnType<typeof dependency_2.org.hyperledger.cacti.plugin.ledger.connector.besu.Web3TransactionReceiptPB.prototype.toObject>;
                callOutput?: ReturnType<typeof dependency_1.google.protobuf.Any.prototype.toObject>;
                success?: boolean;
            } = {};
            if (this.transactionReceipt != null) {
                data.transactionReceipt = this.transactionReceipt.toObject();
            }
            if (this.callOutput != null) {
                data.callOutput = this.callOutput.toObject();
            }
            if (this.success != null) {
                data.success = this.success;
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (this.has_transactionReceipt)
                writer.writeMessage(472834426, this.transactionReceipt, () => this.transactionReceipt.serialize(writer));
            if (this.has_callOutput)
                writer.writeMessage(401260801, this.callOutput, () => this.callOutput.serialize(writer));
            if (this.success != false)
                writer.writeBool(256557056, this.success);
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): InvokeContractV1ResponsePB {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new InvokeContractV1ResponsePB();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 472834426:
                        reader.readMessage(message.transactionReceipt, () => message.transactionReceipt = dependency_2.org.hyperledger.cacti.plugin.ledger.connector.besu.Web3TransactionReceiptPB.deserialize(reader));
                        break;
                    case 401260801:
                        reader.readMessage(message.callOutput, () => message.callOutput = dependency_1.google.protobuf.Any.deserialize(reader));
                        break;
                    case 256557056:
                        message.success = reader.readBool();
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): InvokeContractV1ResponsePB {
            return InvokeContractV1ResponsePB.deserialize(bytes);
        }
    }
}
