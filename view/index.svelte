<script>
  export let posts = []
</script>
<svelte:head>
  <script src="https://cdn.tailwindcss.com"></script>
</svelte:head>



<div class="px-3 md:px-7 lg:px-10 py-7">
  <div class="text-center">
    <h1 class="text-4xl text-center">Hacker News Clone ~ Bud</h1>
    <p>This is just an Tailwind Overlay of Live bud demo, Hacker News.
      <a class="text-cyan-900 hover:text-purple-700 hover:font-bold underline" href="https://github.com/livebud/bud">Bud</a> is a Full-Stack Web Framework for Go.</p>
  </div>
  <div class="grid place-items-center">
    <div id="posts" class="h-auto text-sm text-sm columns-1 sm:columns-1 md:columns-2 lg:columns-3 xl:columns-4 2xl:columns-5 2xl:text-md gap-3 mx-auto space-y-9 py-9 md:px-3">
      {#each posts as post}
        <div class="break-inside-avoid" id="{post.id}" title="{post.id}">
          <a href="{post.url}" class="block max-w-sm rounded-xl mx-2 hover:bg-blue-100 cursor-pointer hover:scale-110 duration-300 overflow-hidden shadow-lg">
            <div class="px-6 py-4">
              <div class="font-bold text-3xl mb-2"><strong class="block font-medium mb-1">{post.title}</div>
              <div>
                <span class="text-neutral-500 inline-block" title="">
                  {post.created_at}
                </span>
                    <span class="text-neutral-500 inline-block" title="">
                  <b>{post.num_comments}</b> engagements
                </span>
                <span class="p-0.5 rounded bg-teal-50 text-neutral-500 inline-block" title="">
                  <b>{post.points}</b> points
                </span>
              </div>
              <div class="text-neutral-500" title="">
                Posted by <strong>{post.author}</strong>
              </div>
            </div>
          </a>
        </div>
      {/each}
    </div>
  </div>
</div>

<style>
  table {
    border-collapse: collapse;
  }
</style>
