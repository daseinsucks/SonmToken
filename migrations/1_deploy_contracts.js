const USDTReceiver = artifacts.require("USDTReceiver");
const SonmToken = artifacts.require("SonmToken");

module.exports = async function (deployer, network, accounts) {
  const usdtAddress = "0xYourUSDTContractAddress"; // Replace with your actual USDT contract address
  const minter = accounts[0]; // Address that will be used as the minter for SonmToken

  // Deploy USDTReceiver
  await deployer.deploy(USDTReceiver, usdtAddress);
  const usdtReceiverInstance = await USDTReceiver.deployed();

  // Deploy SonmToken with the address of the USDTReceiver as the minter
  const tokenName = "SonmToken";
  const tokenTicker = "SONM";
  await deployer.deploy(SonmToken, usdtReceiverInstance.address, tokenName, tokenTicker);
  const sonmTokenInstance = await SonmToken.deployed();

  // Set the SonmToken address in the USDTReceiver contract
  await usdtReceiverInstance.setToken(sonmTokenInstance.address);

  console.log("Deployment completed successfully!");
  console.log("USDTReceiver Address:", usdtReceiverInstance.address);
  console.log("SonmToken Address:", sonmTokenInstance.address);
};