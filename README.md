
# HildaFetch 
Um sistema de *fetch* configurável para exibir informações do sistema com arte ASCII, cores e bordas personalizadas.
Feito com Golang, Gopsutil e Lipgloss 

---

## 🎨 Personalização da Arte ASCII

O conteúdo do arquivo:

```
template/ascii.txt
```

será exibido exatamente como está, no lado esquerdo do *fetch*.

Você pode usar qualquer arte ASCII.
Exemplo:

```
  /\_/\  
 ( o.o )  
  > ^ <
```
> AVISO: fique atento ao tamanho da arte ascii, asciis muito grandes podem quebrar a visualização do fetch.

---

## 🛠️ Configuração do Template de Informações

O arquivo:

```
template/template.txt
```

define o layout e a ordem das informações.

### Formato:

```
*Texto Livre* : {{PLACEHOLDER}}
```

* O texto entre ** é exibido exatamente como escrito (você pode incluir emojis, símbolos, etc.).
* O valor dentro de `{{ }}` será substituído pela informação correspondente.
* O : é o separador e você pode mudar isso também.
### Placeholders disponíveis:

* `{{CPU}}` — Nome do processador
* `{{OS}}` — Sistema operacional
* `{{RAM}}` — Memória RAM
* `{{Username}}` — Nome do usuário
* `{{Uptime}}` — Tempo de atividade

---

## 🎨 Configuração de Estilo e Cores

As opções de estilo estão em:

```
config.json
```

Exemplo:

```json
{
    "username_color": "#80B9B9",
    "colorPrimary": "#3B6E8F",
    "colorSecundary": "#D46A6A",
    "borderColor": "#7BA670",
    "borderType": "rounded",

    "frases": [
        "Well, that was traumatic... but that’s the life of an adventurer.",
        "The wilderness is calling.",
        "Woffs are just misunderstood."
    ]
}
```

### Significado das opções:

| Campo            | Função                                                                              |
| ---------------- | ----------------------------------------------------------------------------------- |
| `username_color` | Cor do nome do computador                                                           |
| `colorPrimary`   | Cor do texto base e da ASCII                                                        |
| `colorSecundary` | Cor das informações do sistema                                                      |
| `borderColor`    | Cor da borda                                                                        |
| `borderType`     | Estilo da borda: `normal`, `rounded`, `thick`, `hidden`, `ascii`, `double`, `block` |
| `frases`         | Lista de frases aleatórias a serem exibidas                                         |

---

## ✅ Uso Básico

1. Defina sua arte ASCII em `template/ascii.txt`
2. Configure as informações desejadas em `template/template.txt`
3. Ajuste as cores e bordas em `config.json`
4. Execute o programa normalmente

---
