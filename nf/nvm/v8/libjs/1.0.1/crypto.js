// Copyright (C) 2018 go-nebulas authors
// 
// This file is part of the go-nebulas library.
// 
// the go-nebulas library is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
// 
// the go-nebulas library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
// 
// You should have received a copy of the GNU General Public License
// along with the go-nebulas library.  If not, see <http://www.gnu.org/licenses/>.
// 

'use strict';

var Crypto = function() {
    Object.defineProperty(this, "nativeCrypto", {
        configurable: false,
        enumerable: false,
        get: function(){
            return _native_crypto;
        }
    });
};

Crypto.prototype = {
 
    sha256: function(data) {
        if (typeof data !== "string") {
            throw new Error("input must be string");
        }
        return this.nativeCrypto.sha256(data);
    },

    sha3256: function(data) {
        if (typeof data !== "string") {
            throw new Error("input must be string");
        }
        return this.nativeCrypto.sha3256(data);
    },

    ripemd160: function(data) {
        if (typeof data !== "string") {
            throw new Error("input must be string");
        }
        return this.nativeCrypto.ripemd160(data);
    },

    recoverAddress: function(alg, data, sign) {
        if (!Number.isInteger(alg)) {
            throw new Error("alg must be integer");
        }

        if (typeof data != "string" || typeof sign != "string") {
            throw new Error("data & sign must be string");
        }

        return this.nativeCrypto.recoverAddress(alg, data, sign);
    },
};

module.exports = new Crypto();