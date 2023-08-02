# LavaVM

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

LavaVM is an eduicational project aimed at building a Java Virtual Machine (JVM) implementation in Go.

## Components

- **Class Parser**: Responsible for parsing .class files. It reads and validates the magic byte, minor and major version, and the constant pool count.
- **Constant Pool Parser**: Reads constant pool entries from the .class file. Currently supports parsing UTF8, Integer, Float, Long, Double, Class, String, FieldRef, MethodRef, InterfaceMethodRef, NameAndType, MethodHandle, MethodType, Dynamic, InvokeDynamic, and Module constants.
- **Execution Engine**: Finds the main mentod, reads the bytecode, then starts executing it

# References

- [Java Class File Format](https://en.wikipedia.org/wiki/Java_class_file)
- [Java Class File JVM Spec](https://docs.oracle.com/javase/specs/jvms/se6/html/ClassFile.doc.html)

## License

LavaVM is licensed under the MIT License. See [LICENSE](LICENSE) for more information.