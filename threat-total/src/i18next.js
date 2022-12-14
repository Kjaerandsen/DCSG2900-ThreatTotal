import i18n from "i18next";
import { initReactI18next } from "react-i18next";
import resources from './translations/translationData.json'
import LanguageDetector from 'i18next-browser-languagedetector'

i18n
  .use(initReactI18next) // passes i18n down to react-i18next
  .use(LanguageDetector)
  .init({
    resources,
    ns: ['translation', 'translation2', 'about'],
    fallbackLng: "en", // default language
    debug: false,
    detector: {
      order: ['queryString', 'localStorage', 'cookie'],
      caches: ['localStorage', 'cookie']
    },
    interpolation: {
        escapeValue: false // Not needed in react
    }
});

export default i18n