// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "./SonmToken.sol";

contract USDTReceiver is Ownable {
    IERC20 public usdt;
    SonmToken public sonmToken;
    address public treasury;
    uint256 public mintPrice;
    uint256 public feePercentage;

    constructor(address _usdt) 
    Ownable(msg.sender)
    {  
        usdt = IERC20(_usdt);
        treasury = msg.sender;
        mintPrice = 100;
        feePercentage = 5;  
    }

    function mintForUSDT(uint256 amount) public {
        require(usdt.transferFrom(msg.sender, address(this), amount), "USDT transfer failed");
        sonmToken.mint(msg.sender, amount*mintPrice); // Minting 1:100 ratio of USDT to SonmToken
    }

     function burnAndSendUSDT(uint256 amount) external {
        sonmToken.burnFrom(msg.sender, amount);

        uint256 usdtAmount = amount / 100;
        uint256 fee = usdtAmount / 100 * 5;
        usdtAmount -= fee;

        require(usdt.transfer(msg.sender, usdtAmount), "USDT transfer failed");
        require(usdt.transfer(treasury, fee), "fee transfer failed");
    }

    function setToken(address _sonmToken) public onlyOwner {
        sonmToken = SonmToken(_sonmToken);
    }

    function setMintPrice(uint256 _newPrice) public onlyOwner {
        mintPrice = _newPrice;
    }
    function setFee(uint256 _newFee) public onlyOwner {
        feePercentage = _newFee;
    }
    function setTreasury(address _newTreasury) public onlyOwner {
        treasury = _newTreasury;
    }
}