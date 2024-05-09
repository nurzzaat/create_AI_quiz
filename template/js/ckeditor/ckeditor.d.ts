/**
 * @license Copyright (c) 2014-2023, CKSource Holding sp. z o.o. All rights reserved.
 * For licensing, see LICENSE.md or https://ckeditor.com/legal/ckeditor-oss-license
 */
import { ClassicEditor } from '@ckeditor/ckeditor5-editor-classic';
import { Alignment } from '@ckeditor/ckeditor5-alignment';
import { Autoformat } from '@ckeditor/ckeditor5-autoformat';
import { Bold, Italic, Underline } from '@ckeditor/ckeditor5-basic-styles';
import { BlockQuote } from '@ckeditor/ckeditor5-block-quote';
import type { EditorConfig } from '@ckeditor/ckeditor5-core';
import { Essentials } from '@ckeditor/ckeditor5-essentials';
import { FontSize } from '@ckeditor/ckeditor5-font';
import { Heading } from '@ckeditor/ckeditor5-heading';
import { HorizontalLine } from '@ckeditor/ckeditor5-horizontal-line';
import { GeneralHtmlSupport } from '@ckeditor/ckeditor5-html-support';
import { Image, ImageCaption, ImageInsert, ImageResize, ImageToolbar, ImageUpload } from '@ckeditor/ckeditor5-image';
import { AutoLink, Link } from '@ckeditor/ckeditor5-link';
import { List, ListProperties } from '@ckeditor/ckeditor5-list';
import { MediaEmbed, MediaEmbedToolbar } from '@ckeditor/ckeditor5-media-embed';
import { Paragraph } from '@ckeditor/ckeditor5-paragraph';
import { PasteFromOffice } from '@ckeditor/ckeditor5-paste-from-office';
import { SpecialCharacters, SpecialCharactersArrows, SpecialCharactersCurrency, SpecialCharactersEssentials } from '@ckeditor/ckeditor5-special-characters';
import { Table, TableProperties, TableToolbar } from '@ckeditor/ckeditor5-table';
import { SimpleUploadAdapter } from '@ckeditor/ckeditor5-upload';
declare class Editor extends ClassicEditor {
    static builtinPlugins: (typeof Alignment | typeof AutoLink | typeof Autoformat | typeof BlockQuote | typeof Bold | typeof Essentials | typeof FontSize | typeof GeneralHtmlSupport | typeof Heading | typeof HorizontalLine | typeof Image | typeof ImageCaption | typeof ImageInsert | typeof ImageResize | typeof ImageToolbar | typeof ImageUpload | typeof Italic | typeof Link | typeof List | typeof ListProperties | typeof MediaEmbed | typeof MediaEmbedToolbar | typeof Paragraph | typeof PasteFromOffice | typeof SimpleUploadAdapter | typeof SpecialCharacters | typeof SpecialCharactersArrows | typeof SpecialCharactersCurrency | typeof SpecialCharactersEssentials | typeof Table | typeof TableProperties | typeof TableToolbar | typeof Underline)[];
    static defaultConfig: EditorConfig;
}
export default Editor;
