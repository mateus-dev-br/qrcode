### QR Code Generator // Gerador de QR Code

### How to Use

### 1. Prerequisites

-   **Go** installed on your machine. If you don't have Go installed, follow [these instructions](https://golang.org/doc/install).

### 2. Clone the Repository

First, clone the repository to your local machine:
```bash
git clone https://github.com/seuusuario/go-qrcode.git 
cd go-qrcode
```
### 3. Install Dependencies

The project depends on the `go-qrcode` library to generate QR Codes. Install the dependencies with the command:

```bash
go get -u github.com/skip2/go-qrcode
```
### 4. Run the Code

Now, you can run the code to generate a QR Code. Execute the following command:
```go
go run main.go
```
This will generate a file named `qrcode.png` in the same directory.

### 5. Modify the QR Code Content

You can modify the content of the QR Code by changing the `content` variable in the `main.go` file. For example:
```css
content := "www.seulink.com"
```
Change it to the link or text you want to encode in the QR Code.

----------

### Code Explanation

-   The code uses the `go-qrcode` library to generate the QR Code.
-   The content of the QR Code is passed in the `content` variable.
-   The generated QR Code will be saved as a PNG file named `qrcode.png`.
-   The code uses the error correction level `qrcode.Medium` and a size of 256x256 pixels.

---

### Como Usar

### 1. Pré-requisitos

-   **Go** instalado na sua máquina. Se não tiver o Go instalado, siga [essas instruções](https://golang.org/doc/install).

### 2. Clonar o Repositório

Primeiro, clone o repositório para a sua máquina local:
Altere para o link ou texto que você deseja codificar no QR Code.
```bash
git clone https://github.com/seuusuario/go-qrcode.git 
cd go-qrcode
```
### 3. Instalar Dependências

O projeto depende da biblioteca `go-qrcode` para gerar os QR Codes. Instale as dependências com o comando:


```bash
go get -u github.com/skip2/go-qrcode
```


### 4. Rodar o Código

Agora você pode rodar o código para gerar um QR Code. Execute o comando:

```go
go run main.go
```
Isso irá gerar um arquivo chamado `qrcode.png` no mesmo diretório.


### 5. Modificar o Conteúdo do QR Code

Você pode modificar o conteúdo do QR Code alterando a variável `content` no arquivo `main.go`. Por exemplo:

```css
content := "www.seulink.com"
```
Altere para o link ou texto que você deseja codificar no QR Code.

### Explicação do Código
-   O código usa a biblioteca `go-qrcode` para gerar o QR Code.
-   O conteúdo do QR Code é passado na variável `content`.
-   O QR Code gerado será salvo em um arquivo PNG chamado `qrcode.png`.
-   O código usa o nível de correção de erro `qrcode.Medium` e uma dimensão de 256x256 pixels.
---
