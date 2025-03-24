package contas

import "banco/clientes"

type ContaCorrente struct {
	Titular 		clientes.Titular 
	NumeroAgencia 	int 
	NumeroConta 	int
	saldo 			float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque >= 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	}

	if valorDoSaque < 0 {
		return "O valor não pode ser menor que 0"
	}
	return "Saldo insuficiente"
}

func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {
	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return "Deposito realizado com sucesso", c.saldo
	}

	return "O valor do depósito é menor que 0", c.saldo
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia <= c.saldo && valorDaTransferencia > 0{
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	}	

	return false 
}

func (c *ContaCorrente) ObterSaldo() float64{
	return c.saldo
}