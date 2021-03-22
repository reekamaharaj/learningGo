
# code-review-golang

This is a Go code review exercise. 

Problem Context: The Nomad scheduler tracks different tasks running on client nodes. They are called "allocations". We want to add an API for showing the top K most resource intensive allocations on each client node, where resources are CPU, disk and memory. 

The code in this PR is for the backend for such an API. It creates a tracker to update allocations by their resource utilization, and methods to return top K by each resource. 

Review the provided pull request, making code review comments as you go. The scope of your comments can be suggestions, design improvements and identifying defects or bugs. You are asked to do what you would normally do in a typical code review.
