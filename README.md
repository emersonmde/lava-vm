# LavaVM

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

LavaVM is a project aimed at building a Java Virtual Machine (JVM) implementation in Go. This project is in early
development and is mostly a mechanism for me to learn Go. 


## Components

- **Class Parser**: Responsible for parsing .class files. It reads and validates the magic byte, minor and major version, and the constant pool count.
- **Constant Pool Parser**: Reads constant pool entries from the .class file. Currently supports parsing UTF8, Integer, Float, Long, Double, Class, String, FieldRef, MethodRef, InterfaceMethodRef, NameAndType, MethodHandle, MethodType, Dynamic, InvokeDynamic, and Module constants.

# References

- [Java Class File Format](https://en.wikipedia.org/wiki/Java_class_file)

## License

LavaVM is licensed under the MIT License. See [LICENSE](LICENSE) for more information.