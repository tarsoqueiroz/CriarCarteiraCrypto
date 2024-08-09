# Criar Carteira para Crypto

## Sobre

Simples rotina em Golang para criação de carteira para cryptomoedas com foco em Bitcoin.

Uma carteira de Bitcoin não é como uma carteira física, ela não armazena suas moedas. Em vez disso, ela armazena suas chaves privadas. Essas chaves são sequências aleatórias de caracteres que dão a você o controle sobre seus Bitcoins.

> **Lembre-se:** A segurança da sua carteira é sua responsabilidade. Ao trabalhar com criptomoedas, sempre priorize a segurança e busque conhecimento aprofundado sobre o assunto.

## Referências

- [Making your own safety cold Ethereum HD wallet using Golang](https://huangwenwei.com/blogs/making-you-own-ethereum-hd-cold-wallet-using-golang)
- [Documentação da btcsuite](https://pkg.go.dev/github.com/btcsuite/btcutil)
- [Documentação da libsecp256k1](https://pkg.go.dev/github.com/nbd-wtf/go-nostr/libsecp256k1)
- [createWallet.js](https://raw.githubusercontent.com/digitalinnovationone/formacao-blockchain-dio/main/Modulo%2001%20Fundamentos%20da%20Blockchain/Curso%2001%20Introducao%20a%20Blockchain/Criando%20e%20utilizando%20a%20sua%20carteira%20de%20criptomoedas/src/createWallet.js)
- [bitcoinjs-lib](https://bitcoinjs.github.io/bitcoinjs-lib/modules/networks.html) (NodeJS)
- [BIP39](https://www.npmjs.com/package/bip39)
- [BIP32]()

## Código

**Componentes Essenciais de uma Carteira**:

- **Chave Privada:** A chave secreta que permite gastar seus Bitcoins.
- **Endereço Público:** Derivado da chave privada, é o endereço para onde você recebe Bitcoins.
- **Mnemônico:** Uma frase de palavras que serve como backup para a sua chave privada.

**Bibliotecas Go para Bitcoin**:

Para criar uma carteira em Go, podemos utilizar bibliotecas como:

- **[btcsuite](https://pkg.go.dev/github.com/btcsuite/btcutil):** Uma suíte completa de pacote Go para interagir com a rede Bitcoin.
- **[libsecp256k1](https://pkg.go.dev/github.com/nbd-wtf/go-nostr/libsecp256k1):** Uma biblioteca C para a curva elíptica secp256k1 usada pelo Bitcoin, com bindings para Go.

**Código Básico para Gerar uma Carteira**:

- [main.go](./src/main.go)

**Explicação do Código**:

- **Geração da Seed:** A função GenerateSeed cria uma sequência aleatória de bytes que servirá como base para a nossa chave privada.
- **Criação da Chave Mestre:** A chave mestre é criada a partir da seed.
- **Derivação de uma Chave:** Derivamos uma chave filha da chave mestre para nossa carteira.
- **Obtenção do Endereço Público:** A partir da chave pública, calculamos o endereço para onde os Bitcoins serão enviados.
- **Mnemônico:** A seed é convertida em um mnemônico, uma frase fácil de lembrar que pode ser usada para restaurar a carteira.

**Considerações Importantes**:

- **Segurança:**
  - **Armazenamento da Chave Privada:** Nunca compartilhe sua chave privada com ninguém.
  - **Backup do Mnemônico:** Guarde seu mnemônico em um lugar seguro e offline.
  - **Hardware Wallet:** Considere usar uma hardware wallet para maior segurança.
- **Rede:** Para interagir com a rede Bitcoin, você precisará de um nó completo ou utilizar um serviço de carteira.
- **Transações:** A criação e o envio de transações envolvem outros conceitos como fees, inputs e outputs.
- **Segurança de Endereços:** Evite reutilizar endereços e utilize endereços gerados hierarquicamente.

Este código fornece uma base sólida para criar uma carteira Bitcoin em Go. Para uma implementação completa, você precisará adicionar:

- **Gerenciamento de múltiplas carteiras:** Crie um sistema para gerenciar várias carteiras.
- **Transações:** Implemente funções para criar e enviar transações.
- **Interface do usuário:** Crie uma interface para interagir com a carteira.
- **Integração com a rede Bitcoin:** Conecte sua carteira à rede Bitcoin para sincronizar e enviar transações.

## Recursos Adicionais

### Electrum Wallet

Electrum is a lightweight Bitcoin wallet.

- Official documentation: [electrum.readthedocs.io](https://electrum.readthedocs.io/)

**Setup**:

