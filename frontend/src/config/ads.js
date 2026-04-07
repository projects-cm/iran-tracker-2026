/**
 * Ad Network Configuration
 *
 * To enable a network, set enabled: true and paste your network's
 * script/iframe HTML into the `html` field.
 *
 * To go live with AADS:
 *   1. Go to https://a-ads.com/
 *   2. Create a campaign unit, get the <iframe> snippet
 *   3. Paste the full <iframe> string as the `html` value below
 *   4. Set enabled: true
 */
export const AD_CONFIG = {
  // Master kill-switch. Set to false to hide all ads instantly.
  enabled: true,

  slots: {
    // Leaderboard below the header (728x90 on desktop, 320x50 on mobile)
    leaderboard: {
      enabled: true,
      label: 'SPONSORED',
      // Replace with your actual AADS or PropellerAds snippet:
      html: null,
      fallbackWidth: 728,
      fallbackHeight: 90,
    },

    // Sidebar rectangle — reserved for future sidebar layout (300x250)
    sidebar: {
      enabled: false,
      label: 'SPONSORED',
      html: null,
      fallbackWidth: 300,
      fallbackHeight: 250,
    },
  },
};
