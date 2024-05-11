const router = require("express").Router();
const Web3 = require('web3');

// Web3 provider 설정
const web3 = new Web3(process.env.API_URI); 

// Petchain 스마트 컨트랙트 ABI 및 주소
// 주의: 아래 ABI와 주소는 예시입니다. 실제 Petchain 스마트 컨트랙트의 ABI와 주소를 사용해야 합니다.
const contractABI = [...]; // Petchain 스마트 컨트랙트의 ABI로 변경
const contractAddress = '0x...'; // Petchain 스마트 컨트랙트의 주소로 변경
const contract = new web3.eth.Contract(contractABI, contractAddress);

// 이더리움 계정 주소
const account = '0x...';

// 반려동물 등록 및 증명서 발급

router.post('/register/pet', async (req, res) =>{
    try {
        const {petName, petAge, petType} = req.body;
        
        // 반려동물 정보를 블록체인에 저장
        // Petchain 스마트 컨트랙트의 반려동물 등록 함수 사용: registerPet()
        // 생성된 반려동물의 고유 아이디 리턴

        res.send({petName, petAge, petType});
    } catch (error) {
        console.log({error});
        res.send({error});
    }
});

// 반려동물 고유 ID를 통하여 Petchain 컨트랙트에 반려동물 증서를 생성
router.post('/issue/certificate', async (req, res) =>{
    try {
        const {petId, petName, petAge, petType, certificateId} = req.body;
        
        // 반려동물 증서 데이터 등록
        async function registerPetCertificate(petId, petName, petAge, petType, certificateId, issueDate) {
            const registerPet = contract.methods.registerPet(petId, petName, petAge, petType);
            const issueCertificate = contract.methods.issueCertificate(petId, certificateId, issueDate);
        
            // 반려동물 등록 및 증서 발행 트랜잭션 전송
            const gasPrice = await web3.eth.getGasPrice();
            const gasEstimate1 = await registerPet.estimateGas({ from: account });
            const gasEstimate2 = await issueCertificate.estimateGas({ from: account });
        
            const registerTx = await registerPet.send({ from: account, gas: gasEstimate1, gasPrice });
            const issueTx = await issueCertificate.send({ from: account, gas: gasEstimate2, gasPrice });
        
            return { registerTx, issueTx };
        }
        
        const result = await registerPetCertificate(petId, petName, petAge, petType, certificateId, Date.now());

        res.send({petId, result});
    } catch (error) {
        console.log({error});
        res.send({error});
    }
});

// 반려동물 증서 사용
router.post('/use/certificate', (req, res) =>{
    try {
        const {certificateId, petId} = req.body;
        
        // Petchain 컨트랙트에서 반려동물 증서를 사용
        // Petchain 스마트 컨트랙트의 증서 사용 함수 사용: useCertificate()
        // 위 함수를 사용하면 내부적으로 verifyCertificate와 invalidateCertificate가 실행됨

        res.send({petId});
    } catch (error) {
        console.log({error});
        res.send({error});
    }
});

module.exports = router;
