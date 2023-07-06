from datasets import load_dataset
import os
from pathlib import Path

# Set the cache directory to the same directory as this script
os.environ["HF_DATASETS_CACHE"] = str(Path(__file__).parent)


dataset = load_dataset("quora")
