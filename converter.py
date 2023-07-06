from datasets import load_dataset

dataset = load_dataset("quora")

# Export the 'train' split of the dataset to a CSV file
dataset["train"].to_csv("quora_train.csv")
