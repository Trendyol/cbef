function OnUpdate(doc, meta) {
	log("Doc created/updated", meta.id);
}

function OnDelete(meta, options) {
	log("Doc deleted/expired", meta.id);
}
