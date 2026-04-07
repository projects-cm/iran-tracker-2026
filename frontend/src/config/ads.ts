import type { AdConfig } from '../types';

/**
 * Ad Network Configuration
 *
 * To go live with AADS:
 *   1. Sign up at https://a-ads.com/
 *   2. Create a campaign unit, get the <iframe> snippet
 *   3. Paste the full <iframe> string as the `html` value below
 *   4. Ensure `enabled: true`
 */
export const AD_CONFIG: AdConfig = {
  // Master kill-switch — set to false to hide ALL ads instantly
  enabled: true,

  slots: {
    // Leaderboard below the header (728×90 desktop, 320×50 mobile)
    leaderboard: {
      enabled: true,
      label: 'SPONSORED',
      html: null, // ← paste your AADS/PropellerAds <iframe> here
      fallbackWidth: 728,
      fallbackHeight: 90,
    },

    // Sidebar rectangle — reserved for future sidebar layout (300×250)
    sidebar: {
      enabled: false,
      label: 'SPONSORED',
      html: null,
      fallbackWidth: 300,
      fallbackHeight: 250,
    },
  },
};
