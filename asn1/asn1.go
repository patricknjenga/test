package asn1

import `encoding/asn1`

type NRT0201 struct {
	// Define the fields of the NRT0201 struct according to the ASN.1 schema
	NRTRDE NRT0201NRTRDE
}

type NRT0201NRTRDE struct {
	Version    asn1.RawValue
	FileType   asn1.RawValue
	FileFormat asn1.RawValue
	File       []NRT0201File
}

type NRT0201File struct {
	FileHeader  NRT0201FileHeader
	FileContent asn1.RawValue
}

type NRT0201FileHeader struct {
	FileName   asn1.RawValue
	FileType   asn1.RawValue
	FileFormat asn1.RawValue
	FileSize   asn1.RawValue
	RecordSize asn1.RawValue
	Record     []NRT0201Record
}

type NRT0201Record struct {
	RecordHeader NRT0201RecordHeader
	RecordData   NRT0201RecordData
}

type NRT0201RecordHeader struct {
	RecordType asn1.RawValue
	RecordSize asn1.RawValue
}

type NRT0201RecordData struct {
	RecordData asn1.RawValue
}
