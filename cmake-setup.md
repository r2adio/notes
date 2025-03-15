# CMake:

## Minimal CMake projects:

- content of **CMakeLists.txt**
  ```txt
  cmake_minimum_required(Version 3.16..3.24)
  project(our_project Version 0.0.1
                      Description "Our first project"
                      LANGUAGE CXX)
  ```

## Using CMake:
- Older method:
  ```txt
  cd <project_dir>
  mkdir build
  cd build
  cmake ..  # configure build and generate a Makefile
  make -j8  # run make with 8 threads
  ```

- Latest method:
  - for reference [check](https://stackoverflow.com/questions/67425557/how-do-i-build-a-cmake-project) 
  ```txt
  cmake -DCMAKE_BUILD_TYPE=Release -S /path/to/source-dir -B /path/to/build-dir
  cmake --build /path/to/build-dir
  ```
