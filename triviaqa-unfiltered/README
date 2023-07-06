-------------------------------------------------------------------------------------------------------
The University of Washington TriviaQA Dataset (version unfiltered 1.0)
-------------------------------------------------------------------------------------------------------

TriviaQA is a reading comprehension dataset. This release contains the unfiltered version with 110K question-answer pairs and many more documents than the dataset described in the paper. We only provide Wikipedia and Web documents from the top 10 search results. However, full documents can be crawled using the URLs referenced in the data. Not all documents for a given question contain the answer string(s). We expect this to make the QA task much harder.



@InProceedings{JoshiTriviaQA2017,
  author    = {Joshi, Mandar  and  Choi, Eunsol  and  Weld, Daniel S. and Zettlemoyer, Luke},
  title     = {TriviaQA: A Large Scale Distantly Supervised Challenge Dataset for Reading Comprehension},
  booktitle = {Proceedings of the 55th Annual Meeting of the Association for Computational Linguistics (Volume 1: Long Papers)},
  month     = {July},
  year      = {2017},
  address   = {Vancouver, Canada},
  publisher = {Association for Computational Linguistics},
}

-------------------------------------------------------------------------------------------------------
LIST OF DATA FILES
* unfiltered-train.json, unfiltered-dev.json
These files contain (verified) questions, answers, and document names in the train/dev set for the wikipedia/web domain. The details of the verified evaluation set are described in section 4 (evidence analysis) of the paper. The wikipedia/web documents are listed in the json array "EntityPages"/"SearchResults" for each question. The "Filename" field in each element of the array indicates the relative path of the file inside the wikipedia/web directory.

* unfiltered-test-without-answers.json
This file contains questions and document names (no answers) in the unfiltered test set. We are withholding the test answers for a later release. Please check the website -- http://nlp.cs.washington.edu/triviaqa/ -- for details of the test evaluation. 

The evidence documents have been packaged in the RC version of the dataset. 

-------------------------------------------------------------------------------------------------------
CODE
Code, including evaluation scripts, can be found at
https://github.com/mandarjoshi90/triviaqa

-------------------------------------------------------------------------------------------------------
CONTACT
Mandar Joshi 			mandar90@cs.washington.edu
Eunsol Choi   			eunsol@cs.washington.edu
Dan Weld       			weld@cs.washington.edu
Luke Zettlemoyer 		lsz@cs.washington.edu

-------------------------------------------------------------------------------------------------------
LEGAL
The University of Washington does not own the copyright of the questions and documents included in TriviaQA. We are not responsible for their content or meaning. The dataset is intended for non-commercial research purposes only. By downloading the dataset, you agree not to reproduce, duplicate, copy, sell, trade, resell or exploit for any commercial purpose, any portion of the contexts and any portion of derived data. 

