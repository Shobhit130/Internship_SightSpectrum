// #include <iostream>
// #include<climits>
// using namespace std;
// int main(){
//     int n;
//     cin>>n;
//     int a[n];
//     for(int i=0;i<n;i++){
//         cin>>a[i];
//     }
//     //N can vary from 1 to 10^6
//     const int N = 1e6+2; //10^6+2
//     int idx[N];
//     for(int i=0;i<N;i++){
//         idx[i] = -1;
//     }
//     int minidx = INT_MAX;

//     for(int i=0;i<n;i++){
//         if(idx[a[i]] != -1){
//             if(minidx>idx[a[i]]){
//                 minidx = idx[a[i]];
//             }
//         }else{
//             idx[a[i]] = i;
//         }
//     }
//     cout<<minidx+1<<endl;
// }
// #include<iostream>
// using namespace std;
// int main(){
//     int n,s;
//     cin>>n>>s;
//     int a[n];
//     for(int i=0;i<n;i++){
//         cin>>a[i];
//     }
//     int st=0,en=0,cursum=0;
//     while(cursum+a[en]<=s){
//         cursum += a[en];
//         en++;
//     }
//     cout<<cursum<<" "<<en<<endl;
//     cursum += a[en];
//     while(cursum>s){
//         cursum -= a[st];
//         st++;
//     }
//     if(st>en){
//         cout<<-1<<endl;
//     }else{
//         cout<<st<<" "<<en<<endl;
//     }
// }
#include<iostream>
using namespace std;
int main(){
    int n,s;
    cin>>n>>s;
    int arr[n];
    for(int i=0;i<n;i++){
        cin>>arr[i];
    }
    int curr_sum = arr[0], start = 0, i;
     for (i = 1; i <= n; i++) {
        while (curr_sum > s && start < i - 1) {
            curr_sum = curr_sum - arr[start];
            start++;
        }
        if (curr_sum == s) {
            cout << "Sum found between indexes "
                 << start << " and " << i - 1;
                 return 0;
     }
      if (i < n)
            curr_sum = curr_sum + arr[i];
    }
    cout<<"No subarray"<<endl;
}