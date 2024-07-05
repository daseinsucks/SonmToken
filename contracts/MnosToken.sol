// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Permit.sol";

contract MnosToken is ERC20, ERC20Burnable, Ownable, ERC20Permit {
   uint8 public constant DECIMALS = 6;
    constructor(address minter, string memory name, string memory ticker)
        ERC20(name, ticker)
        Ownable(msg.sender)
        ERC20Permit(name)
    {
        minterContract = minter;
    }
address public minterContract;

    modifier onlyMinter {
        require(msg.sender == minterContract);
        _;
    }

    function mint(address to, uint256 amount) public onlyMinter {
        _mint(to, amount);
    }

    function ownerMint(address to, uint256 amount) public onlyOwner {
        _mint(to, amount);
    }
     function decimals() public view virtual override returns (uint8) {
        return DECIMALS;
    }
}

