package payments

type PaymentMethod interface {
	Pay(usd int) int
	Cancel(id int)
}

type PaymentModel struct {
	payments      map[int]PaymentInfo
	paymentMethod PaymentMethod
}

func NewPaymentModel(paymentMethod PaymentMethod) *PaymentModel {
	return &PaymentModel{
		payments:      make(map[int]PaymentInfo),
		paymentMethod: paymentMethod,
	}
}

func (p *PaymentModel) Pay(desc string, usd int) int {
	id := p.paymentMethod.Pay(usd)
	info := PaymentInfo{
		Description: desc,
		USD:         usd,
		Cancelled:   false,
	}
	p.payments[id] = info
	return id
}

func (p *PaymentModel) Cancel(id int) {
	info, ok := p.payments[id]
	if !ok {
		return
	}
	p.paymentMethod.Cancel(id)
	info.Cancelled = true
	p.payments[id] = info
}

func (p *PaymentModel) Info(id int) PaymentInfo {
	info, ok := p.payments[id]
	if !ok {
		return PaymentInfo{}
	}
	return info
}

func (p *PaymentModel) AllInfo() map[int]PaymentInfo {
	tempmap := make(map[int]PaymentInfo, len(p.payments))
	for k, v := range p.payments {
		tempmap[k] = v
	}
	return tempmap
}
