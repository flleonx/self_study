import { Blog } from "./blog.js";

async function main() {
  const blog = new Blog();

  await blog.initialize();

  const posts = await blog.getAllPost();

  if (posts.length === 0) {
    console.log(
      "No post avaliable. Run `node import-post.js` to load some sample posts"
    );
  }

  for (const post of posts) {
    console.log(post.title);
    console.log('-'.repeat(post.title.length));
    console.log(`Published on ${new Date(post.created_at).toISOString()}`);
    console.log(post.content);
  }
}

main().catch(console.error);
