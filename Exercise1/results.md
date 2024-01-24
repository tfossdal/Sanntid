Sharing a variabel:

Both the C ang Go code failed to return 0, due to both accessing the same variable while the other one edits it, and so it will not always be properly updated.

Limiting the CPU to using one core returns the correct answer, now there is now problems to paralell processing (since nothing happens in paralell anymore).

