const HDWalletProvider = require('@truffle/hdwallet-provider');
require('dotenv').config();

// 환경 변수에서 필요한 정보를 불러옵니다.
const { API_URI, PRIVATE_KEY, MNEMONIC, API_KEY, ETHERSCAN_API_KEY } = process.env;

module.exports = {
    development: {
        host: '127.0.0.1',
        port: 8545,
        network_id: '*', // 어떤 네트워크든 연결 가능
    },
    networks: {
        goerli: {
            provider: () => new HDWalletProvider(PRIVATE_KEY.toString(), API_URI),
            network_id: 5, // Goerli 테스트넷의 네트워크 ID
            gasPrice: 20000000000, // 20 GWei
            gas: 6000000, // 컨트랙트 배포 시 예상 가스 한도
        },
    },
    compilers: {
        solc: {
            version: '0.8.4', // solc 버전 지정
        },
    },
    plugins: ['truffle-plugin-verify'],
    api_keys: {
      etherscan: ETHERSCAN_API_KEY, // Etherscan API 키를 여기에 지정합니다. 아직 키가 없다면 공란으로 두세요.
    }
};
