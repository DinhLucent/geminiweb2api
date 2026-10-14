import json
import time
import requests

BASE = "http://localhost:4981/gemini/v1beta"

with requests.post(
    f"{BASE}/interactions",
    json={
        "input": "Compare Elasticsearch vs OpenSearch architecture and performance",
        "agent": "deep-research-pro-preview-12-2025",
        "background": True,
        "stream": True       # Enable streaming
    },
    stream=True,
    headers={"Accept": "text/event-stream"}
) as resp:
    for line in resp.iter_lines(decode_unicode=True):
        if not line or not line.startswith("data:"):
            continue
        event = json.loads(line[5:].strip())

        status = event.get("status", "")

        if status == "in_progress":
            print(f"[...] Researching...", end="\r", flush=True)

        elif status == "completed":
            outputs = event.get("outputs") or []
            if outputs:
                print("\n" + outputs[-1]["text"])
            break

        elif status == "failed":
            print("Error:", event.get("error"))
            break