<script setup>
const { issue, pr, commit } = defineProps(["issue", "pr", "commit"]);
</script>

<template>
  (<span v-if="issue"
    ><a :href="'https://github.com/summit-fi/goverter/issues/' + issue"
      >#{{ issue }}
    </a>
    via </span
  ><a v-if="pr" :href="'https://github.com/summit-fi/goverter/pull/' + pr"
    >#{{ pr }}</a
  ><a v-if="commit" :href="'https://github.com/summit-fi/goverter/commit/' + commit"
    >{{ commit.slice(0, 6) }}</a
  >)
</template>
