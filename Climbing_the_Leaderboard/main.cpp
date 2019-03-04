// #include <bits/stdc++.h>
#include <iostream>
#include <vector>
#include <string>
#include <fstream>
using namespace std;

vector<string> split_string(string);

struct Item {
  Item(int a, int b): index(a), value(b) {}
  int index;
  int value;
};

// Complete the climbingLeaderboard function below.
vector<int> climbingLeaderboard(vector<int> scores, vector<int> alice) {
  vector<Item> list;
  vector<int> ans;
  int lastValue = -1;
  int idx = 1;
  int l = scores.size();
  for (int i = 0; i < l; i++) {
    if (i == 0) {
      lastValue = scores[i];
      list.push_back(Item{idx, scores[i]});
      continue;
    }  
    // unique 
    if (scores[i] < lastValue) {
      idx++;
      list.push_back(Item{idx, scores[i]});
      lastValue = scores[i];
    }
  }
  int cursor = list.size() - 1;
  for(auto i: alice) {
    for(; cursor >= 0 && list[cursor].value < i; cursor--);
    if (cursor < 0) {
      ans.push_back(1);
      continue;
    }
    if (i == list[cursor].value) {
      ans.push_back(list[cursor].index);
    } else if (i < list[cursor].value){
      ans.push_back(list[cursor].index + 1);
    }
  }
  return ans;
}


int main()
{
    ofstream fout(getenv("OUTPUT_PATH"));

    int scores_count;
    cin >> scores_count;
    cin.ignore(numeric_limits<streamsize>::max(), '\n');

    string scores_temp_temp;
    getline(cin, scores_temp_temp);

    vector<string> scores_temp = split_string(scores_temp_temp);

    vector<int> scores(scores_count);

    for (int i = 0; i < scores_count; i++) {
        int scores_item = stoi(scores_temp[i]);

        scores[i] = scores_item;
    }

    int alice_count;
    cin >> alice_count;
    cin.ignore(numeric_limits<streamsize>::max(), '\n');

    string alice_temp_temp;
    getline(cin, alice_temp_temp);

    vector<string> alice_temp = split_string(alice_temp_temp);

    vector<int> alice(alice_count);

    for (int i = 0; i < alice_count; i++) {
        int alice_item = stoi(alice_temp[i]);

        alice[i] = alice_item;
    }

    vector<int> result = climbingLeaderboard(scores, alice);

    for (int i = 0; i < result.size(); i++) {
        fout << result[i];

        if (i != result.size() - 1) {
            fout << "\n";
        }
    }

    fout << "\n";

    fout.close();

    return 0;
}

vector<string> split_string(string input_string) {
    string::iterator new_end = unique(input_string.begin(), input_string.end(), [] (const char &x, const char &y) {
        return x == y and x == ' ';
    });

    input_string.erase(new_end, input_string.end());

    while (input_string[input_string.length() - 1] == ' ') {
        input_string.pop_back();
    }

    vector<string> splits;
    char delimiter = ' ';

    size_t i = 0;
    size_t pos = input_string.find(delimiter);

    while (pos != string::npos) {
        splits.push_back(input_string.substr(i, pos - i));

        i = pos + 1;
        pos = input_string.find(delimiter, i);
    }

    splits.push_back(input_string.substr(i, min(pos, input_string.length()) - i + 1));

    return splits;
}
