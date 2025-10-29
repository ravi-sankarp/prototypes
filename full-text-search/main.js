// Build a tiny in-memory inverted index
const docs = {
  1: 'laptop case',
  2: 'lapdop bag',
  3: 'desktop stand',
  4: 'lapto cooling pad'
};

function tokenize(text) {
  return text.toLowerCase().split(/\W+/).filter(Boolean);
}

// Build term → docIDs map
const invertedIndex = {};
for (const [docId, text] of Object.entries(docs)) {
  for (const term of tokenize(text)) {
    if (!invertedIndex[term]) invertedIndex[term] = new Set();
    invertedIndex[term].add(docId);
  }
}

console.log('Inverted Index:', invertedIndex);

function levenshtein(a, b) {
  const dp = Array.from({ length: a.length + 1 }, () =>
    Array(b.length + 1).fill(0)
  );

  for (let i = 0; i <= a.length; i++) dp[i][0] = i;
  for (let j = 0; j <= b.length; j++) dp[0][j] = j;

  for (let i = 1; i <= a.length; i++) {
    for (let j = 1; j <= b.length; j++) {
      const cost = a[i - 1] === b[j - 1] ? 0 : 1;
      dp[i][j] = Math.min(
        dp[i - 1][j] + 1, // deletion
        dp[i][j - 1] + 1, // insertion
        dp[i - 1][j - 1] + cost // substitution
      );
    }
  }

  return dp[a.length][b.length];
}

function fuzzySearch(query, maxEdits = 1) {
  const results = [];
  for (const term of Object.keys(invertedIndex)) {
    const distance = levenshtein(query, term);
    if (distance <= maxEdits) {
      results.push({ term, distance, docs: [...invertedIndex[term]] });
    }
  }
  return results.sort((a, b) => a.distance - b.distance);
}

function searchDocs(query) {
  const matches = fuzzySearch(query);
  const docIds = new Set(matches.flatMap((m) => m.docs));
  return [...docIds].map((id) => ({ id, text: docs[id] }));
}

console.log("Documents matching 'laptop~1':");
console.log(searchDocs('laptop'));
