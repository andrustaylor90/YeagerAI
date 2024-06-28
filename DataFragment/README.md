# Main Code

## simpleHash Function:

- This function generates a 30-character hash from a given string using SHA-256.
- It then truncates the resulting hash to 30 characters.

## reconstructData Function:

- This function takes a map of fragments, where each fragment has a sequence number (key) and a dictionary containing the data and its hash.
- It sorts the keys to ensure the fragments are processed in the correct order.
- It then verifies the integrity of each fragment by comparing the provided hash with the hash computed by the `simpleHash` function.
- If all fragments pass the verification, it concatenates their data to reconstruct the original string.
- If any fragment fails the verification, it returns an error message.

# Testing Code

## TestSimpleHash:

- This test verifies that the simpleHash function produces a 30-character hash.
- It also checks that different inputs produce different hashes and the same input produces the same hash.

## TestReconstructData:

- This test checks if the reconstructData function correctly reconstructs the original data from valid fragments.

## TestReconstructDataWithInvalidHash:

- This test verifies that the reconstructData function returns an error message when a fragment has an invalid hash.

## TestReconstructDataWithOutOfOrderFragments:

- This test checks if the reconstructData function correctly handles out-of-order fragments.