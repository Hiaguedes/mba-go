# Stack

- stack trabalha de forma segmentada
    vantagem: eficiencia de memoria e melhoria de desempenho (aumento/ diminuicao de tamanho)
    inicializacao da stack com apenas 2kb
    conforme a goroutine vai sendo executada, o runtime verifica se ha necessidade de mais espaco na stack (checkpoints)
    e alocado um novo segmento de stack, os dados da stack antiga sao copiadas para o novo segmento e os ponteiros atualizados. (normalmente o tamanho do novo segmento e o dobro do anterior)

- reducao de tamanho
    se uma goroutine precisou de uma stack maior mas agora ela nao possui mais essa necessidade, o runtime pode ,eventualmente, liberar os segmentos que nao sao mais necessarios

- guard pages e stack overflow
    o go usa "guard pages" para proteger contra acessos fora dos limites da stack, prevenindo falhas de segmentacao (segmentation faults)

    verifica limites das stack pra garantir que nao cresca ate um determinado ponto para evitar stack overflow

### O que fica na stack

- variaveis locais
- parametros de funcoes
- enderecos de retorno
- escopo e tempo de vida

## Como saber se algo foi pra stack ou pra heap?

- decisao tomada pelo compilador
- "escape analysis"
    se uma variavel e usada de uma forma que ela pode sobreviver alem do escopo da funcao onde ela foi criada, ela e movida pra heap

    para verificar se algo escapou pra heap, utilize: go build -gcflags="-m" main.go

    pra pegar de funcoes recursivas voce usa o debugging dlv

### Escape analysis

- retorno de ponteiros
    se uma funcao retorna um ponteiro para uma variavel local, essa variavel deve ser alocada na hap, pois o ponteiro pode ser usado fora do escopo da funcao

- armazenamento em estrutura de dados
    se uma variavel local e armazenada em uma estrutura de dados que sobrevive ao escopo da funcao, ela deve ser alocada na heap

- goroutines
    se uma variavel local e usada dentro de uma goroutine, ela deve ser alocada na heap, pois a goroutine pode continuar a execucao apos a funcao retornar
