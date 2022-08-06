# lattigo learning
A library for lattice-based homomorphic encryption in Go
[original repo](https://github.com/tuneinsight/lattigo)
this repo just for learning purposes.

Lattigo：基于格的多方同态加密库,具备如下特性：

全RNS BFV和CKKS方案及其各自多方版本的实现。
性能可与最先进的C++库相媲美。
密集键和稀疏键高效、高精度的自举过程，适用于全RNS CKKS。
支持跨平台构建的纯 Go 实现，包括浏览器客户端的 WASM 编译。

相关名词释义

格：指Lattigo是基于格密码方案实现的，格是一种数学结构，本质上由N维空间中的一组点组成，并具有一定的周期结构，格密码方案是指基于基向量来生成最短的独立向量作为秘密向量（数学家和密码学家们普遍认为，对于一个维数足够高的格，通过一组随机选取的格基找到一组短格基，或得到一组线性无关的短格向量是困难的。这个问题称作最短独立向量问题）。

BFV和CKKS：两种全同态加密的两种方案。
RNS：剩余数系统，用较少的数表示较多的数。
源码结构

lattigo/ring：RNS基中多项式的模块化算术运算。
lattigo/bfv：提供对整数的模块化算术。
lattigo/ckks：在复数和实数上提供近似算术。
lattigo/dbfv和lattigo/dckks：BFV和CKKS方案的多方版本，支持具有秘密共享密钥的安全多方计算解决方案。
lattigo/rlwe和lattigo/drlwe：基于 RLWE 的通用多方同态加密的通用基础。

Lattigo方案的可行性分析
使用lattigo/bfv来作为全同态加密合约的支撑工具，使得智能合约失去了浮点数计算能力。本质上来讲，对精度要求严苛的场景都不应使用浮点数，浮点数应通过同比扩大或者同比缩小来消除小数或还原小数。同时lattigo/bfv提供加运算、减运算、非运算、模运算、乘运算、除运算等，满足基本运算需求。Fabric的合约具有完整且独立的运行环境，理论上任意的应用程序在接入合约API后都可以作为合约运行，满足lattigo的运行需求。
