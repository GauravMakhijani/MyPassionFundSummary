package literals

const (
	ErrInvalidInput  = "err- invalid input format"
	ErrMarshal       = "err- marshalling ping response"
	ErrCreatingPDF   = "error generating pdf file"
	PDFFormat        = "P"
	ExcelFormat      = "E"
	ErrInvalidFormat = "err- invalid format"
	DateFormat       = "02/01/2006"
)

type PageStyle string


const	PAGEORIENTATION  PageStyle = "L"
const	UNITSTR         PageStyle = "mm"
const	PAGESIZE        PageStyle = "A4"

