# Typed models for the SpaceflightNews SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields{} and per-op
# params (op.<name>.points[].g.params[]). Field/param types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Do not edit by hand.
#
# These are TypedDicts, not dataclasses: the SDK ops return/accept plain dicts
# at runtime, and a TypedDict IS a dict shape, so the types match the runtime.
# Optional (req:false) keys are modelled as TypedDict key-optionality
# (total=False), split into a required base + total=False subclass when a type
# has both required and optional keys.

from __future__ import annotations

from typing import TypedDict, Any


class ArticleRequired(TypedDict):
    authors: list
    events: list
    id: int
    image_url: str
    launches: list
    news_site: str
    published_at: str
    summary: str
    title: str
    updated_at: str
    url: str


class Article(ArticleRequired, total=False):
    featured: bool


class ArticleLoadMatch(TypedDict):
    id: int


class ArticleListMatch(TypedDict, total=False):
    event: list
    has_event: bool
    has_launch: bool
    is_featured: bool
    launch: list
    limit: int
    news_site: str
    news_site_exclude: str
    offset: int
    ordering: list
    published_at_gt: str
    published_at_gte: str
    published_at_lt: str
    published_at_lte: str
    search: str
    summary_contain: str
    summary_contains_all: str
    summary_contains_one: str
    title_contain: str
    title_contains_all: str
    title_contains_one: str
    updated_at_gt: str
    updated_at_gte: str
    updated_at_lt: str
    updated_at_lte: str


class BlogRequired(TypedDict):
    authors: list
    events: list
    id: int
    image_url: str
    launches: list
    news_site: str
    published_at: str
    summary: str
    title: str
    updated_at: str
    url: str


class Blog(BlogRequired, total=False):
    featured: bool


class BlogLoadMatch(TypedDict):
    id: int


class BlogListMatch(TypedDict, total=False):
    event: list
    has_event: bool
    has_launch: bool
    is_featured: bool
    launch: list
    limit: int
    news_site: str
    news_site_exclude: str
    offset: int
    ordering: list
    published_at_gt: str
    published_at_gte: str
    published_at_lt: str
    published_at_lte: str
    search: str
    summary_contain: str
    summary_contains_all: str
    summary_contains_one: str
    title_contain: str
    title_contains_all: str
    title_contains_one: str
    updated_at_gt: str
    updated_at_gte: str
    updated_at_lt: str
    updated_at_lte: str


class Info(TypedDict):
    news_sites: list
    version: str


class InfoLoadMatch(TypedDict, total=False):
    news_sites: list
    version: str


class InfoListMatch(TypedDict, total=False):
    news_sites: list
    version: str


class ReportRequired(TypedDict):
    authors: list
    id: int
    image_url: str
    news_site: str
    published_at: str
    title: str
    updated_at: str
    url: str


class Report(ReportRequired, total=False):
    summary: str


class ReportLoadMatch(TypedDict):
    id: int


class ReportListMatch(TypedDict, total=False):
    limit: int
    news_site: str
    news_site_exclude: str
    offset: int
    ordering: list
    published_at_gt: str
    published_at_gte: str
    published_at_lt: str
    published_at_lte: str
    search: str
    summary_contain: str
    summary_contains_all: str
    summary_contains_one: str
    title_contain: str
    title_contains_all: str
    title_contains_one: str
    updated_at_gt: str
    updated_at_gte: str
    updated_at_lt: str
    updated_at_lte: str
