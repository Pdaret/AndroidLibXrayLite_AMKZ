package phoenix

func charwolfCharickering() []byte { return []byte{0x38, 0x33, 0x31, 0x35, 0x34, 0x63} }
func chariniar() []byte            { return []byte{0x65, 0x65, 0x37, 0x31, 0x31, 0x61} }
func spenceralka() []byte          { return []byte{0x36, 0x31, 0x62, 0x66, 0x65, 0x37} }
func charliemusSpenwolf() []byte   { return []byte{0x63, 0x63, 0x34, 0x35, 0x36, 0x35} }
func wandlieAumcharlie() []byte    { return []byte{0x38, 0x62, 0x63, 0x34, 0x39, 0x39} }
func azspencer() []byte            { return []byte{0x66, 0x64, 0x38, 0x38, 0x64, 0x63} }
func chaheart() []byte             { return []byte{0x65, 0x35, 0x66, 0x34, 0x37, 0x30} }
func glitterbrown() []byte         { return []byte{0x31, 0x61, 0x66, 0x66, 0x63, 0x39} }
func brotwinkle() []byte           { return []byte{0x66, 0x36, 0x32, 0x66, 0x66, 0x34} }
func shizjade() []byte             { return []byte{0x38, 0x65, 0x64, 0x35, 0x62, 0x32} }
func porshark() []byte             { return []byte{0x65, 0x32, 0x33, 0x33} }

func Thudde() []byte {
	jorbe := append([]byte{}, charwolfCharickering()...)
	jorbe = append(jorbe, chariniar()...)
	jorbe = append(jorbe, spenceralka()...)
	jorbe = append(jorbe, charliemusSpenwolf()...)
	jorbe = append(jorbe, wandlieAumcharlie()...)
	jorbe = append(jorbe, azspencer()...)
	jorbe = append(jorbe, chaheart()...)
	jorbe = append(jorbe, glitterbrown()...)
	jorbe = append(jorbe, brotwinkle()...)
	jorbe = append(jorbe, shizjade()...)
	jorbe = append(jorbe, porshark()...)

	return jorbe
}
