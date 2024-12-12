from pathlib import Path
import json

cwd = Path(__file__).parent

fileaddr = cwd / "webScraping/main_scarped_1.json"

cc = {}
with open(fileaddr,"r") as  filer:
    cc = json.load(filer)


print(cc["Proposals"]["Writing Your First Proposal to the ArbitrumDAO"]["content"])