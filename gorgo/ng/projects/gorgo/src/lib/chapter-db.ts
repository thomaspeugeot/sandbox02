 // generated by GenNgClass.go

import { ParagraphDB} from './paragraph-db'

export class ChapterDB {
	CreatedAt?: string;
	DeletedAt?: string;
	ID?: number;

	Name?: string
	Title?: string


	Paragraphs?: Array<ParagraphDB>

	// ID generated for the implementation of the field Chapter{}.Chapters []*Chapter
	Chapter_ChaptersDBID?: number

	// ID generated for the implementation of the field Chapter{}.Chapters []*Document
	Document_ChaptersDBID?: number

}
