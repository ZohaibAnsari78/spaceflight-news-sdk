# SpaceflightNews SDK utility: make_context

from spaceflightnews_sdk.core.context import SpaceflightNewsContext


def make_context_util(ctxmap, basectx):
    return SpaceflightNewsContext(ctxmap, basectx)
