pragma solidity ^0.5.5;
pragma experimental ABIEncoderV2;

contract Sample {
    uint256 val;

    function setValue(uint256 _val) public {
        val = _val;
    }

    function getValue() public view returns (uint256) {
        return val;
    }

    event A(address indexed val_0, address indexed val_1);

    function setterA(address val_0, address val_1) public payable {
        emit A(val_0, val_1);
    }

    function setA1() public payable {
        emit A(
            0x0000000000000000000000000000000000000000,
            0x0100000000000000000000000000000000000000
        );
    }

    function setA2() public payable {
        emit A(
            0x0100000000000000000000000000000000000000,
            0x0000000000000000000000000000000000000000
        );
    }
}
