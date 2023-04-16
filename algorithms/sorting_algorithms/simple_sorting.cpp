#include <cstdlib>
#include <iostream>
#define MAX 100
using namespace std;

int main()
{
  int n;
  int arr[MAX];
  srand(time(0));

  cout << "Enter a number N: " << endl;
  cin >> n;

  // inputing values in an array
  for (int i = 0; i < n; i++) {
    arr[i] = rand();
  }

  // outputting the unsorted array
  for (int i = 0; i < n; i++) {
    cout << arr[i] << " ";
  }

  cout << endl;

  // sorting and array
  for (int i = 0; i < n; i++) {
    for (int j = i + 1; j < n; j++) {
      // ascending order
      if (arr[j] < arr[i]) {
        int temp = arr[i];
        arr[i] = arr[j];
        arr[j] = temp;
      }
    } 
  }

  // display the sorted array
  cout << "sorted" << endl;
  for (int i = 0; i < n; i++) {
    cout << arr[i] << " ";
  }

  cout << endl;
  
  return 0;
}
