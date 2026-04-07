import React, { useRef } from 'react';
import { AD_CONFIG } from '../config/ads';

interface BannerAdProps {
  slot: string;
}

export default function BannerAd({ slot }: BannerAdProps) {
  const containerRef = useRef<HTMLDivElement>(null);

  const slotConfig = AD_CONFIG?.slots?.[slot];

  if (!AD_CONFIG.enabled || !slotConfig?.enabled) return null;

  const { html, label, fallbackWidth, fallbackHeight } = slotConfig;

  return (
    <div
      className="banner-ad-wrapper"
      style={{
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
        width: '100%',
        background: 'rgba(15, 23, 42, 0.8)',
        borderBottom: '1px solid rgba(30, 41, 59, 0.8)',
        padding: '6px 0 4px',
      }}
    >
      <span
        style={{
          fontSize: '9px',
          letterSpacing: '0.15em',
          color: '#475569',
          fontFamily: 'monospace',
          marginBottom: '4px',
          textTransform: 'uppercase',
        }}
      >
        {label}
      </span>

      {html ? (
        <div
          ref={containerRef}
          dangerouslySetInnerHTML={{ __html: html }}
          style={{ maxWidth: fallbackWidth, width: '100%' }}
        />
      ) : (
        <div
          style={{
            width: Math.min(fallbackWidth, window.innerWidth - 32),
            height: fallbackHeight,
            background: 'rgba(30, 41, 59, 0.5)',
            border: '1px dashed rgba(71, 85, 105, 0.6)',
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center',
            color: '#334155',
            fontSize: '12px',
            fontFamily: 'monospace',
            borderRadius: '4px',
          }}
        >
          {`[ AD SLOT: ${slot} — ${fallbackWidth}×${fallbackHeight} ]`}
        </div>
      )}
    </div>
  );
}
