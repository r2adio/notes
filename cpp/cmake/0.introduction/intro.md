# CMake:

## 1. definition:

- tool taht generates native build system files.
- making porject portable across different systems and envirnoments.

## • nede of cmake:

- Cross_Platform Build System

  - can generate platform-specific build files, from a single configuration file (CMakeLists.txt).

- simplifies build system maintenance

  - writing and maintaining platform specific build scripts is tedious and error prone.

- Dependency Management

  - help manage dependencies and link libraries easily.

- Complex Projects

  - large projects, with multiple source files and dependencies, CMake simplifies the process of defining and managing build configurations.

- out-of-source builds

  - cmake encourages a clean separation of source code and build artifacts by supporting "out-of-source builds".
  - this avoid polluting the source directory with build files.

- Integration with IDEs

  - modern IDEs have build-in support for CMake.

- Support for modern c++ features
  - simplifies enabling modern c++ standards, by allowing you to set compiler flag easily, using cmd like:
  ```cpp
  set(CMAKE_CXX_STANDARD 21)
  ```

## working of cmake:

1. define build rules in CMakeLists.txt

```bash
cmake_minimum_required(VERSION 3.10)
project(MyProject)
# specify the executable
add_executable(MyExecutable main.cpp)
```

2. configure the build system

```bash
mkdir build
cd build
cmake ..
```

3. build the project

```bash
make
```

![hello](/Users/gaurav/Downloads/blue.png)
