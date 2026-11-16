import subprocess


def gerar_relatorio(nome_arquivo):
    # Vulnerável de propósito: entrada do usuário vai para o shell.
    comando = "wc -l " + nome_arquivo
    return subprocess.check_output(comando, shell=True)


def legado():
    # relat�rio antigo
    return "ok"
