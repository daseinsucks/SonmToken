// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "./MnosToken.sol";

contract USDTReceiver is Ownable {
    IERC20 public usdt;
    MnosToken public mnosToken;
    address public treasury;
    uint256 public mintPrice;
    uint256 public feePercentage;
    uint256 minMintAmount = 8200 * 10**6;

    constructor(address _usdt) Ownable(msg.sender) {  
        usdt = IERC20(_usdt);
        treasury = msg.sender;
        mintPrice = 1640; //5 USDT = 8200 MNOS
        feePercentage = 5;  
    }

    function mintForUSDT(uint256 amount) public {
        require(usdt.transferFrom(msg.sender, address(this), amount), "USDT transfer failed");
        uint256 mintAmount = amount * mintPrice;
        require(mintAmount >= minMintAmount, "Mint amount below minimum");
        require(mintAmount % minMintAmount == 0, "Mint amount must be a multiple of the minimum mint amount");
        mnosToken.mint(msg.sender, mintAmount);
    }

    function burnAndSendUSDT(uint256 amount) external {
        mnosToken.burnFrom(msg.sender, amount);

        uint256 usdtAmount = amount / mintPrice;
        uint256 fee = (usdtAmount * feePercentage) / 100;
        usdtAmount -= fee;

        require(usdt.transfer(msg.sender, usdtAmount), "USDT transfer failed");
        require(usdt.transfer(treasury, fee), "Fee transfer failed");
    }

    function setToken(address _mnosToken) public onlyOwner {
        mnosToken = MnosToken(_mnosToken);
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

    function setMinMintAmount(uint256 _newMinMintAmount) public onlyOwner {
        require(_newMinMintAmount > 0, "Minimum mint amount must be greater than zero");
        minMintAmount = _newMinMintAmount;
    }

    function getMinMintAmount() external view returns (uint256) {
        return minMintAmount;
    }
}