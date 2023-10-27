import bittensor

def query(context: str):
    out = ""
    while out == "":
        out = bittensor.prompt(context)
    return out