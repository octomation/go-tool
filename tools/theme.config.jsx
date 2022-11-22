export default {
  project: {
    link: 'https://github.com/octomation/go-tool',
  },

  docsRepositoryBase: 'https://github.com/octomation/go-tool/blob/main/tools',
  feedback: {
    useLink() {
      return 'https://github.com/octomation/go-tool/discussions/new/choose'
    },
  },
  useNextSeoProps() {
    return {
      titleTemplate: '%s',
    }
  },

  head: (
    <>
      <meta charSet="utf-8"/>
      <meta name="viewport" content="width=device-width, initial-scale=1.0"/>

      <style>{`
        main p a img { display: inline; } /* badges */
      `}</style>
    </>
  ),
  logo: (
    <>
      <img width={24} height={24}
           src="https://raw.githubusercontent.com/octomation/.github/main/.static/octolab.png"
           alt="OctoLab"
      />
      <span>Tool</span>
    </>
  ),
  banner: {
    text: <a href="https://github.com/octomation/go-tool/releases/tag/v1.0.0" target="_blank">
      🎉 Tool v1.0.0 is released. Read more →
    </a>,
  },
  footer: {
    text: <span>
      MIT {new Date().getFullYear()} © <a href="https://github.com/octolab" target="_blank">OctoLab</a>.
    </span>,
  },
}
