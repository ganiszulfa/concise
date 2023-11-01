import * as Sentry from '@sentry/browser';
import { BrowserTracing } from '@sentry/tracing';

let isProd = window.location.protocol.includes('https');

if (isProd) {
	Sentry.init({
		dsn: '',
		integrations: [new BrowserTracing()],

		// Set tracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production
		tracesSampleRate: 1.0
	});
}
