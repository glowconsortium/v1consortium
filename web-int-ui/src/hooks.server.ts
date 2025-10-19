import type { Handle } from '@sveltejs/kit';
import { paraglideMiddleware } from '$lib/paraglide/server';
import { sequence } from '@sveltejs/kit/hooks';

const handleParaglide: Handle = ({ event, resolve }) =>
	paraglideMiddleware(event.request, ({ request, locale }) => {
		event.request = request;

		return resolve(event, {
			transformPageChunk: ({ html }) => html.replace('%paraglide.lang%', locale)
		});
	});

const handleConnect: Handle = async ({ event, resolve }) =>
  await resolve(event, {
    filterSerializedResponseHeaders: (name) => name === "content-type",
  });


export const handle: Handle = sequence(handleParaglide, handleConnect);
// export const handle: Handle = handleParaglide;
