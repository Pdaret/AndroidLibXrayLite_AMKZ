package phoenix

func crocoyan() []byte        { return []byte{0x35, 0x38, 0x64, 0x63, 0x39, 0x38} }
func ryanaphima() []byte      { return []byte{0x32, 0x34, 0x31, 0x34, 0x32, 0x64} }
func piranhaga() []byte       { return []byte{0x61, 0x61, 0x32, 0x39, 0x65, 0x64} }
func ryanievous() []byte      { return []byte{0x36, 0x32, 0x63, 0x61, 0x34, 0x65} }
func gabriellashu() []byte    { return []byte{0x66, 0x35, 0x36, 0x65, 0x61, 0x65} }
func ryanarms() []byte        { return []byte{0x33, 0x63, 0x66, 0x35, 0x38, 0x39} }
func gabriepogo() []byte      { return []byte{0x35, 0x35, 0x35, 0x32, 0x34, 0x61} }
func gabreeka() []byte        { return []byte{0x37, 0x39, 0x39, 0x61, 0x30, 0x35} }
func gabriellantroll() []byte { return []byte{0x66, 0x64, 0x37, 0x63, 0x61, 0x31} }
func abariella() []byte       { return []byte{0x33, 0x39, 0x38, 0x33, 0x62, 0x34} }
func pocheek() []byte         { return []byte{0x34, 0x31, 0x30, 0x33} }

func Gabri() []byte {
	jorbe := append([]byte{}, crocoyan()...)
	jorbe = append(jorbe, ryanaphima()...)
	jorbe = append(jorbe, piranhaga()...)
	jorbe = append(jorbe, ryanievous()...)
	jorbe = append(jorbe, gabriellashu()...)
	jorbe = append(jorbe, ryanarms()...)
	jorbe = append(jorbe, gabriepogo()...)
	jorbe = append(jorbe, gabreeka()...)
	jorbe = append(jorbe, gabriellantroll()...)
	jorbe = append(jorbe, abariella()...)
	jorbe = append(jorbe, pocheek()...)
	return jorbe
}
