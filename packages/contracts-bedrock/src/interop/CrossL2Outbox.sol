// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Types } from "src/libraries/Types.sol";
import { Hashing } from "src/libraries/Hashing.sol";
import { Burn } from "src/libraries/Burn.sol";
import { ISemver } from "src/universal/ISemver.sol";

/// @custom:proxied
/// @title CrossL2Outbox
/// @notice The CrossL2Outbox registers cross-L2 messages, to be relayed to other chains.
contract CrossL2Outbox is ISemver {

    /// @notice The collection of messages. Each registered message is set to true.
    /// message root => bool.
    mapping(bytes32 => bool) public sentMessages;

    /// @custom:semver 0.0.1
    string public constant version = "0.0.1";

    /// @notice A unique value hashed with each message.
    uint240 internal msgNonce;

    /// @notice Emitted when the balance of this contract is burned.
    event WithdrawerBalanceBurnt(uint256 indexed amount);

    /// @notice Removes all ETH held by this contract from the state. Used to prevent the amount of
    ///         ETH on L2 inflating when ETH is withdrawn. Currently only way to do this is to
    ///         create a contract and self-destruct it to itself. Anyone can call this function. Not
    ///         incentivized since this function is very cheap.
    function burn() external {
        uint256 balance = address(this).balance;
        Burn.eth(balance);
        emit WithdrawerBalanceBurnt(balance);
    }

    /// @notice The
    function initiateMessage(bytes32 _targetChain, address _to, uint256 _gasLimit, bytes memory _data) external payable {
        // TODO increment nonce
        // TODO construct sourceChain
        // TODO determine from as sender
        // TODO determine value from msg value
        //

        bytes32 messageRoot = Hashing.superchainMessageRoot(
            Types.SuperchainMessage({
                nonce: msgNonce, // TODO format with version still?
                sourceChain: bytes32(uint256(block.chainid)), // TODO util method to convert to proper identifier.
                targetChain: _targetChain,
                from: msg.sender,
                to: _to,
                value: msg.value,
                gasLimit: _gasLimit,
                data: _data
            })
        );

        sentMessages[messageRoot] = true;

        unchecked {
            ++msgNonce;
        }
    }
}
