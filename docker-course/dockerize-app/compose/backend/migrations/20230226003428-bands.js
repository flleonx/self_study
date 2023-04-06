export const up = async (db) => {
  await db.collection("albums").insertMany([{ artist: "The Beatles" }]);
};

export const down = async (db) => {
  await db.collection("albums").deleteMany([{ artist: "The Beatles" }]);
};
