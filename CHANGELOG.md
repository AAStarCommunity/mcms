# @smartcontractkit/mcms

## 0.21.0

### Minor Changes

- [#413](https://github.com/smartcontractkit/mcms/pull/413) [`fe08b88`](https://github.com/smartcontractkit/mcms/commit/fe08b8829cfd7818e04b808f08e5c9d0290b5da8) Thanks [@ecPablo](https://github.com/ecPablo)! - Add labels field to proposal base type

- [#402](https://github.com/smartcontractkit/mcms/pull/402) [`6d53434`](https://github.com/smartcontractkit/mcms/commit/6d53434bdce48425dbdaa47a221a9aabd467a9e6) Thanks [@ecPablo](https://github.com/ecPablo)! - Add AllSigners() function to Config struct.

- [#404](https://github.com/smartcontractkit/mcms/pull/404) [`0d1d460`](https://github.com/smartcontractkit/mcms/commit/0d1d4607524fb8a47e3190e8c78811064a58fa68) Thanks [@ecPablo](https://github.com/ecPablo)! - upgrade golangci lint

### Patch Changes

- [#403](https://github.com/smartcontractkit/mcms/pull/403) [`1dbe8a3`](https://github.com/smartcontractkit/mcms/commit/1dbe8a38aa6b31dbdd2838dbf5d9f850da85320c) Thanks [@gustavogama-cll](https://github.com/gustavogama-cll)! - chore(solana): validate AccessController config multiple accounts

- [#409](https://github.com/smartcontractkit/mcms/pull/409) [`2d2b9ea`](https://github.com/smartcontractkit/mcms/commit/2d2b9ea5c7784a9f49aa5bd399fc860fe58b9f4b) Thanks [@akuzni2](https://github.com/akuzni2)! - Add a getter method to get signer address

## 0.20.2

### Patch Changes

- [#396](https://github.com/smartcontractkit/mcms/pull/396) [`a6ff022`](https://github.com/smartcontractkit/mcms/commit/a6ff022f82a2b10eb7ece0b6507a3a43421e0db1) Thanks [@gustavogama-cll](https://github.com/gustavogama-cll)! - fix(solana): fix calculation of AppendInstructionData chunk size

- [#378](https://github.com/smartcontractkit/mcms/pull/378) [`571f37b`](https://github.com/smartcontractkit/mcms/commit/571f37bc207b5ece48413b1f764fb7a2f00bd7b9) Thanks [@gustavogama-cll](https://github.com/gustavogama-cll)! - fix(solana): check current root before calling SetRoot

## 0.20.1

### Patch Changes

- [#385](https://github.com/smartcontractkit/mcms/pull/385) [`4c4a585`](https://github.com/smartcontractkit/mcms/commit/4c4a585eab19cc743b2eb2fb7149f3312d3a85f8) Thanks [@gustavogama-cll](https://github.com/gustavogama-cll)! - chore: cache compiled solana programs in CI

- [#392](https://github.com/smartcontractkit/mcms/pull/392) [`b7948a2`](https://github.com/smartcontractkit/mcms/commit/b7948a228ae1faf333d314c0dfb8ce8e1756fcf4) Thanks [@gustavogama-cll](https://github.com/gustavogama-cll)! - fix: set CU limit as the maximum (1.4M) to avoid CU errors in mainnet

- [#394](https://github.com/smartcontractkit/mcms/pull/394) [`6ff9d02`](https://github.com/smartcontractkit/mcms/commit/6ff9d024cedb9ac81093442407239d08992405ec) Thanks [@kylesmartin](https://github.com/kylesmartin)! - Support Ledger MMII encoding

## 0.20.0

### Minor Changes

- [#388](https://github.com/smartcontractkit/mcms/pull/388) [`8f66eb2`](https://github.com/smartcontractkit/mcms/commit/8f66eb299b2f3c8901fd4e8688565ce79a72a9f0) Thanks [@gustavogama-cll](https://github.com/gustavogama-cll)! - fix(solana): use solanacommon.SendTransaction() that supports retries

## 0.19.3

### Patch Changes

- [#384](https://github.com/smartcontractkit/mcms/pull/384) [`80e82ad`](https://github.com/smartcontractkit/mcms/commit/80e82ad071080c870e876d098eba3827a8bcd4ed) Thanks [@gustavogama-cll](https://github.com/gustavogama-cll)! - fix(solana): use smaller batch "append instruction" chunk size

- [#381](https://github.com/smartcontractkit/mcms/pull/381) [`0b3000a`](https://github.com/smartcontractkit/mcms/commit/0b3000ac7471503e93836a4ecd23b4378555758e) Thanks [@gustavogama-cll](https://github.com/gustavogama-cll)! - fix(solana): use smaller batch size for the "appendSigner" instructions

## 0.19.2

### Patch Changes

- [#376](https://github.com/smartcontractkit/mcms/pull/376) [`cb1e96a`](https://github.com/smartcontractkit/mcms/commit/cb1e96aeeea47a32d24cabf3b67a2d21013f538e) Thanks [@gustavogama-cll](https://github.com/gustavogama-cll)! - chore: bump chainlink-ccip version

## 0.19.1

### Patch Changes

- [#374](https://github.com/smartcontractkit/mcms/pull/374) [`17381e4`](https://github.com/smartcontractkit/mcms/commit/17381e45d25c8690bee810720024cfb2ea67251b) Thanks [@gustavogama-cll](https://github.com/gustavogama-cll)! - fix: call ClearSignatures in Solana's SetRoot upon receiving an "account in use" error

## 0.19.0

### Minor Changes

- [#368](https://github.com/smartcontractkit/mcms/pull/368) [`4e53d4d`](https://github.com/smartcontractkit/mcms/commit/4e53d4db82fba9e0907ae3bda6a7adc2dbc6c026) Thanks [@gustavogama-cll](https://github.com/gustavogama-cll)! - feat: add IsOperationPending() and IsChainPending() to TimelockExecutable

## 0.18.0

### Minor Changes

- [#366](https://github.com/smartcontractkit/mcms/pull/366) [`637d057`](https://github.com/smartcontractkit/mcms/commit/637d0578ec656f4eacbfd06f88a01d9bcada5034) Thanks [@gustavogama-cll](https://github.com/gustavogama-cll)! - fix(solana): fix the accounts added to the BypassExecuteBatch instruction

## 0.17.0

### Minor Changes

- [#355](https://github.com/smartcontractkit/mcms/pull/355) [`28b56e3`](https://github.com/smartcontractkit/mcms/commit/28b56e31d312458457faf12c9e6354018c2c84b2) Thanks [@friedemannf](https://github.com/friedemannf)! - Add Timelock support for Aptos

- [#265](https://github.com/smartcontractkit/mcms/pull/265) [`a51bc1b`](https://github.com/smartcontractkit/mcms/commit/a51bc1bba20e98f14f0973550f90f80fe58dd2a7) Thanks [@friedemannf](https://github.com/friedemannf)! - Add Aptos support

### Patch Changes

- [#350](https://github.com/smartcontractkit/mcms/pull/350) [`46b795b`](https://github.com/smartcontractkit/mcms/commit/46b795bfbb96e826d3c8a6a28e8e1134e71e354d) Thanks [@ecPablo](https://github.com/ecPablo)! - update ledger product ids

## 0.16.1

### Patch Changes

- [#346](https://github.com/smartcontractkit/mcms/pull/346) [`741fc94`](https://github.com/smartcontractkit/mcms/commit/741fc94709beee3bcc15e9c251cc90cd6b82a109) Thanks [@gustavogama-cll](https://github.com/gustavogama-cll)! - fix: use GetLatestBlockhash to retrive curr. block time in Solana

## 0.16.0

### Minor Changes

- [#342](https://github.com/smartcontractkit/mcms/pull/342) [`58c4da1`](https://github.com/smartcontractkit/mcms/commit/58c4da1ae9c47ef84ab72587e78a5bf3d5993a19) Thanks [@ecPablo](https://github.com/ecPablo)! - Add proposal cancel and bypass helpers

## 0.15.0

### Minor Changes

- [#337](https://github.com/smartcontractkit/mcms/pull/337) [`3965c2d`](https://github.com/smartcontractkit/mcms/commit/3965c2d6d3dba9ef86b3e7917ff9d5227bffc4a5) Thanks [@gustavogama-cll](https://github.com/gustavogama-cll)! - feat: add IsOperationDone and IsChainDone to TimelockExecutable

## 0.14.0

### Minor Changes

- [#309](https://github.com/smartcontractkit/mcms/pull/309) [`d6e3b1d`](https://github.com/smartcontractkit/mcms/commit/d6e3b1d21ab50f31276450f2d4e048c90bb7f126) Thanks [@akhilchainani](https://github.com/akhilchainani)! - Add proposal-level decoding related APIs

- [#335](https://github.com/smartcontractkit/mcms/pull/335) [`8b5981a`](https://github.com/smartcontractkit/mcms/commit/8b5981aa999afb6c42f5f81c0a9e4bc0be563e49) Thanks [@gustavogama-cll](https://github.com/gustavogama-cll)! - fix: add context to NewTimelockExecutable()

## 0.13.0

### Minor Changes

- [#329](https://github.com/smartcontractkit/mcms/pull/329) [`031adfb`](https://github.com/smartcontractkit/mcms/commit/031adfb1038d1716e7620104c5af6ff6b807ce4d) Thanks [@graham-chainlink](https://github.com/graham-chainlink)! - feat(solana): override authority on Configurer

## 0.12.3

### Patch Changes

- [#323](https://github.com/smartcontractkit/mcms/pull/323) [`008aec7`](https://github.com/smartcontractkit/mcms/commit/008aec734962e517b5bb9f9cb384c5ac2d08fff8) Thanks [@graham-chainlink](https://github.com/graham-chainlink)! - solana: update rootSignaturesPDA to include authority as new seed

- [#324](https://github.com/smartcontractkit/mcms/pull/324) [`fc04e76`](https://github.com/smartcontractkit/mcms/commit/fc04e76deac968504b9af51c7c0862fff809cfd9) Thanks [@ecPablo](https://github.com/ecPablo)! - allow empty accounts in solana additional fields for transactions

## 0.12.2

### Patch Changes

- [#320](https://github.com/smartcontractkit/mcms/pull/320) [`5383734`](https://github.com/smartcontractkit/mcms/commit/53837346adefec47eb0ab92937029f645baaf5a2) Thanks [@ecPablo](https://github.com/ecPablo)! - downgrade chainlink solana

## 0.12.1

### Patch Changes

- [#318](https://github.com/smartcontractkit/mcms/pull/318) [`e19a319`](https://github.com/smartcontractkit/mcms/commit/e19a3192493fd79e6cf79154da5668bb637c254d) Thanks [@akhilchainani](https://github.com/akhilchainani)! - setPredecessors in constructor and refactor

## 0.12.0

### Minor Changes

- [#297](https://github.com/smartcontractkit/mcms/pull/297) [`2a4a443`](https://github.com/smartcontractkit/mcms/commit/2a4a443552b4098fd6fb23055b8ec5a5fabb110b) Thanks [@graham-chainlink](https://github.com/graham-chainlink)! - fix: make predecessors parameter optional on NewProposal and NewTimelockProposal

- [#310](https://github.com/smartcontractkit/mcms/pull/310) [`d5d9b16`](https://github.com/smartcontractkit/mcms/commit/d5d9b166d63330dae68e518e8dcdd0b4744bada9) Thanks [@gustavogama-cll](https://github.com/gustavogama-cll)! - feat: enable calling SetConfig without sending the transaction

- [#301](https://github.com/smartcontractkit/mcms/pull/301) [`d8a156c`](https://github.com/smartcontractkit/mcms/commit/d8a156c94f693bfeaa63b85a21e47ec4b4f1190e) Thanks [@akhilchainani](https://github.com/akhilchainani)! - Expose additional functions to help with proposal decoding

- [#313](https://github.com/smartcontractkit/mcms/pull/313) [`4a4a5b5`](https://github.com/smartcontractkit/mcms/commit/4a4a5b52d56bdc8b1fc810a11dee58bd588d672d) Thanks [@ecPablo](https://github.com/ecPablo)! - Remove dependency of timelock converter in solana with rpc.Client

### Patch Changes

- [#311](https://github.com/smartcontractkit/mcms/pull/311) [`3808de9`](https://github.com/smartcontractkit/mcms/commit/3808de9556cd10fb713b009fa29b5ed859d9ae65) Thanks [@graham-chainlink](https://github.com/graham-chainlink)! - Update to latest chainlink-ccip/chains/solana

## 0.11.0

### Minor Changes

- [#295](https://github.com/smartcontractkit/mcms/pull/295) [`29afd9d`](https://github.com/smartcontractkit/mcms/commit/29afd9da94f27be613771f393139d3cb8f1d7ab1) Thanks [@akhilchainani](https://github.com/akhilchainani)! - Add additional isReady helpers to API

## 0.10.1

### Patch Changes

- [#291](https://github.com/smartcontractkit/mcms/pull/291) [`f76ce79`](https://github.com/smartcontractkit/mcms/commit/f76ce79b354e3f50fa00e1d7653990013f6a2955) Thanks [@graham-chainlink](https://github.com/graham-chainlink)! - retrigger publishing due to previous failure

## 0.10.0

### Minor Changes

- [#289](https://github.com/smartcontractkit/mcms/pull/289) [`1eeb56a`](https://github.com/smartcontractkit/mcms/commit/1eeb56aee88fb4e13d155892bdd2ea381c04b61d) Thanks [@graham-chainlink](https://github.com/graham-chainlink)! - Address latest breaking change (new Operation Enums and Chunk based instructions for timelock converter)

## 0.9.0

### Minor Changes

- [#277](https://github.com/smartcontractkit/mcms/pull/277) [`cb6ea80`](https://github.com/smartcontractkit/mcms/commit/cb6ea802d44766ac03043cc1308565b3880f6d7c) Thanks [@graham-chainlink](https://github.com/graham-chainlink)! - feat(solana): simulate executor operation mcms

### Patch Changes

- [#284](https://github.com/smartcontractkit/mcms/pull/284) [`78adad0`](https://github.com/smartcontractkit/mcms/commit/78adad0bb5ad8bbe6e92d757269418504195cf90) Thanks [@graham-chainlink](https://github.com/graham-chainlink)! - fix(deps): downgrade geth 1.14.13 -> 1.14.11

## 0.8.0

### Minor Changes

- [#276](https://github.com/smartcontractkit/mcms/pull/276) [`27b77d5`](https://github.com/smartcontractkit/mcms/commit/27b77d5143e48bafc2cb1d1bac6b75389728adc3) Thanks [@akhilchainani](https://github.com/akhilchainani)! - Update constructors to add predecessor proposals for queuing

- [#254](https://github.com/smartcontractkit/mcms/pull/254) [`aad56bd`](https://github.com/smartcontractkit/mcms/commit/aad56bd4a49f703dd9580a6ba0d25abae573cc95) Thanks [@gustavogama-cll](https://github.com/gustavogama-cll)! - feat(solana): add setRoot simulator

- [#279](https://github.com/smartcontractkit/mcms/pull/279) [`3287f3c`](https://github.com/smartcontractkit/mcms/commit/3287f3cf49636f017ef70e88cad68ae4ba535654) Thanks [@ecPablo](https://github.com/ecPablo)! - Fix bug with multichain timelock execution with predecessors calculation

### Patch Changes

- [#274](https://github.com/smartcontractkit/mcms/pull/274) [`28d52c3`](https://github.com/smartcontractkit/mcms/commit/28d52c329b039b6fc94a9e3394b24c564a2e0d5c) Thanks [@graham-chainlink](https://github.com/graham-chainlink)! - fix(solana): fix simulator side effect bug

## 0.7.0

### Minor Changes

- [#255](https://github.com/smartcontractkit/mcms/pull/255) [`e38816f`](https://github.com/smartcontractkit/mcms/commit/e38816f21105e33692a21775fb0e7dcafbd34b95) Thanks [@ecPablo](https://github.com/ecPablo)! - Add config transformer functionality for solana.

- [#257](https://github.com/smartcontractkit/mcms/pull/257) [`31f1e09`](https://github.com/smartcontractkit/mcms/commit/31f1e0946a6cef6f1943ea62089c911817ca1e0d) Thanks [@akhilchainani](https://github.com/akhilchainani)! - Return generic transaction object instead of just hash

### Patch Changes

- [#270](https://github.com/smartcontractkit/mcms/pull/270) [`d6d880c`](https://github.com/smartcontractkit/mcms/commit/d6d880c3e8588494677252d1beda490f1455ac92) Thanks [@graham-chainlink](https://github.com/graham-chainlink)! - fix: incorporate timelock converter breaking api changes for scheduled operations

- [#264](https://github.com/smartcontractkit/mcms/pull/264) [`8849c73`](https://github.com/smartcontractkit/mcms/commit/8849c73b095b5c3df881f65fadeddc8c599e72db) Thanks [@graham-chainlink](https://github.com/graham-chainlink)! - fix(executor): callproxy as config

- [#216](https://github.com/smartcontractkit/mcms/pull/216) [`a481d17`](https://github.com/smartcontractkit/mcms/commit/a481d174ca83eb11aa6b7b4aff1497ec4fb39da6) Thanks [@gustavogama-cll](https://github.com/gustavogama-cll)! - Fix assignment of `GroupSigners` in `ConfigTransformer.ToConfig()`

- [#268](https://github.com/smartcontractkit/mcms/pull/268) [`d28b0df`](https://github.com/smartcontractkit/mcms/commit/d28b0df6b2dcb9796469dac3387524062c69c383) Thanks [@gustavogama-cll](https://github.com/gustavogama-cll)! - Return the solana.rpc.GetTransactionResult struct as the raw transaction of the Solana SDK.

## 0.6.0

### Minor Changes

- [#256](https://github.com/smartcontractkit/mcms/pull/256) [`45c6a2e`](https://github.com/smartcontractkit/mcms/commit/45c6a2edfa1cd641860dcb3f87c0a32a3bbda636) Thanks [@akhilchainani](https://github.com/akhilchainani)! - Allow callProxy execute capability

- [#245](https://github.com/smartcontractkit/mcms/pull/245) [`7a5944e`](https://github.com/smartcontractkit/mcms/commit/7a5944e1940df6a327ea07eec0c41f0211580133) Thanks [@gustavogama-cll](https://github.com/gustavogama-cll)! - Add context and Converter map TimelockConverter.Convert params

- [#231](https://github.com/smartcontractkit/mcms/pull/231) [`a8447e1`](https://github.com/smartcontractkit/mcms/commit/a8447e1727147fe21ff7f2b8186ceeebeef47a23) Thanks [@graham-chainlink](https://github.com/graham-chainlink)! - feat(solana): timelock inspection - operation statuses check

- [#242](https://github.com/smartcontractkit/mcms/pull/242) [`c610826`](https://github.com/smartcontractkit/mcms/commit/c610826d34826e36eabe91e067a6d194398d6564) Thanks [@gustavogama-cll](https://github.com/gustavogama-cll)! - Add TimelockProposal.Convert for solana

- [#238](https://github.com/smartcontractkit/mcms/pull/238) [`abde70c`](https://github.com/smartcontractkit/mcms/commit/abde70ca21a206e6bf452d54e98805d896cd76b5) Thanks [@graham-chainlink](https://github.com/graham-chainlink)! - feat(solana): implement get roles for timelock inspection

- [#236](https://github.com/smartcontractkit/mcms/pull/236) [`150a1f6`](https://github.com/smartcontractkit/mcms/commit/150a1f6fac2ec450a377ae5e818f5d196a257a9e) Thanks [@graham-chainlink](https://github.com/graham-chainlink)! - Use string for inspector return type

- [#209](https://github.com/smartcontractkit/mcms/pull/209) [`a71dd79`](https://github.com/smartcontractkit/mcms/commit/a71dd79442551bdb757b94263d48b9ef7aa2b3b8) Thanks [@gustavogama-cll](https://github.com/gustavogama-cll)! - Add the `Configurer` component and `SetConfig` call to the Solana SDK.

- [#223](https://github.com/smartcontractkit/mcms/pull/223) [`4adb968`](https://github.com/smartcontractkit/mcms/commit/4adb96870c0a3daac98095656d0fea0753367b0d) Thanks [@gustavogama-cll](https://github.com/gustavogama-cll)! - Add a "context" parameter to all APIs that interact with a blockchain.

- [#227](https://github.com/smartcontractkit/mcms/pull/227) [`21d8809`](https://github.com/smartcontractkit/mcms/commit/21d8809c0a5e8048562804a9c1198367a4eabff1) Thanks [@ecPablo](https://github.com/ecPablo)! - Adds Execute functionality to solana SDK

- [#248](https://github.com/smartcontractkit/mcms/pull/248) [`e153c75`](https://github.com/smartcontractkit/mcms/commit/e153c751aa3f048be66a8687ea5f039c147491d4) Thanks [@ecPablo](https://github.com/ecPablo)! - Timelock execute batch on solana SDK.

- [#211](https://github.com/smartcontractkit/mcms/pull/211) [`be76399`](https://github.com/smartcontractkit/mcms/commit/be76399a414053345b0b6e8e5b1eff951a3efd7e) Thanks [@graham-chainlink](https://github.com/graham-chainlink)! - feat(solana): support get opdata, root and root metadata

### Patch Changes

- [#215](https://github.com/smartcontractkit/mcms/pull/215) [`9f39403`](https://github.com/smartcontractkit/mcms/commit/9f394035272baa4f2fcfb33d626486ba113841d7) Thanks [@gustavogama-cll](https://github.com/gustavogama-cll)! - Add the `Executor` component and `SetRoot` call to the Solana SDK.

- [#259](https://github.com/smartcontractkit/mcms/pull/259) [`a4bc13b`](https://github.com/smartcontractkit/mcms/commit/a4bc13b9f8d58fc29f8acb6475929ee757f3588f) Thanks [@anirudhwarrier](https://github.com/anirudhwarrier)! - usbwallet: fix ledger access for latest firmware and add Ledger Flex

- [#225](https://github.com/smartcontractkit/mcms/pull/225) [`7c9cd3d`](https://github.com/smartcontractkit/mcms/commit/7c9cd3d08c4a04cb1dd596b643976a9b96807149) Thanks [@gustavogama-cll](https://github.com/gustavogama-cll)! - Add PDA finders and ContractAddress parser to the Solana SDK

- [#258](https://github.com/smartcontractkit/mcms/pull/258) [`a11a0ee`](https://github.com/smartcontractkit/mcms/commit/a11a0eea5d7321adfabeea0c42131d80a535e0b3) Thanks [@akhilchainani](https://github.com/akhilchainani)! - non-breaking change to allow a salt override to proposals

- [#228](https://github.com/smartcontractkit/mcms/pull/228) [`b953973`](https://github.com/smartcontractkit/mcms/commit/b953973f62b2c2876f55cb050541a3f990cc1ea7) Thanks [@graham-chainlink](https://github.com/graham-chainlink)! - fix(solana): setProgramID on inspection methods

## 0.5.0

### Minor Changes

- [#214](https://github.com/smartcontractkit/mcms/pull/214) [`04cb547`](https://github.com/smartcontractkit/mcms/commit/04cb5474e8ce890566a4e48739a1d917f245c72f) Thanks [@akhilchainani](https://github.com/akhilchainani)! - Add helper function for timelock proposal execution

- [#210](https://github.com/smartcontractkit/mcms/pull/210) [`8431019`](https://github.com/smartcontractkit/mcms/commit/8431019b9a672edf0c257982d677dcf04897c770) Thanks [@akhilchainani](https://github.com/akhilchainani)! - Implement timelock executor interface + EVM implementation

## 0.4.0

### Minor Changes

- [#183](https://github.com/smartcontractkit/mcms/pull/183) [`6dd7030`](https://github.com/smartcontractkit/mcms/commit/6dd7030d76efaa44e75332421c44b1adc1f31728) Thanks [@akhilchainani](https://github.com/akhilchainani)! - Implement EVM Simulations

### Patch Changes

- [#188](https://github.com/smartcontractkit/mcms/pull/188) [`a8db3c4`](https://github.com/smartcontractkit/mcms/commit/a8db3c4ba39d8067bd96fda915544dd17808d599) Thanks [@ecPablo](https://github.com/ecPablo)! - Fix bug in nonce calculation when multiple proposals were executed

## 0.3.0

### Minor Changes

- [#165](https://github.com/smartcontractkit/mcms/pull/165) [`03682a2`](https://github.com/smartcontractkit/mcms/commit/03682a2772b4771f5af05d1ebd49b0a54e30beaf) Thanks [@ecPablo](https://github.com/ecPablo)! - Add SetConfig support for EVM MCMS contract.

## 0.2.0

### Major Changes

- Refactored public API and internal packages to support multiple chain families.
