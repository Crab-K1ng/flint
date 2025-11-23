package parser

func (p *Parser) error(msg string) {
	p.errors = append(p.errors, msg)
}
