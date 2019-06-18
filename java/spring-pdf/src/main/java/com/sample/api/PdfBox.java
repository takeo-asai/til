package com.sample.api;

import java.io.ByteArrayInputStream;
import java.io.ByteArrayOutputStream;
import java.io.IOException;

import org.apache.pdfbox.pdmodel.PDDocument;
import org.apache.pdfbox.pdmodel.PDPage;
import org.apache.pdfbox.pdmodel.PDPageContentStream;
import org.apache.pdfbox.pdmodel.font.PDFont;
import org.apache.pdfbox.pdmodel.font.PDType1Font;

class PdfBox {
    public static ByteArrayInputStream generate() {
        ByteArrayOutputStream out = new ByteArrayOutputStream();
        try {
            PDDocument document = new PDDocument();
            PDPage blank = new PDPage();
            document.addPage(blank);

            PDPageContentStream contentStream = new PDPageContentStream(document, blank);

            contentStream.beginText();
            contentStream.setFont(PDType1Font.HELVETICA, 10);
            contentStream.drawString("abcdefg");
            contentStream.endText();

            contentStream.close();

            document.save(out);
            document.close();

        } catch (IOException e) {
            e.printStackTrace();
        }

        return new ByteArrayInputStream(out.toByteArray());

    }
}